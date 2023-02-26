package gpio

import (
	"container/heap"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

type watcherAction int

const (
	watcherAdd watcherAction = iota
	watcherRemove
	watcherClose
)

type watcherCmd struct {
	pin    Pin
	action watcherAction
}

// WatcherNotification represents a single pin change The new value of the pin numbered by Pin is Value
// CN 表示单个针脚更改用Pin编号的针脚的新值是Value
type WatcherNotification struct {
	Pin   uint
	Value uint
}

type fdHeap []uintptr

func (h fdHeap) Len() int { return len(h) }

// Less is actually greater (we want a max heap)
// CN 更少实际上更大（我们想要最大堆）
func (h fdHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h fdHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *fdHeap) Push(x interface{}) {
	*h = append(*h, x.(uintptr))
}

func (h *fdHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h fdHeap) FdSet() *syscall.FdSet {
	fdset := &syscall.FdSet{}
	for _, val := range h {
		fdset.Bits[val/64] |= 1 << (uint(val) % 64)
	}
	return fdset
}

const watcherCmdChanLen = 32
const notificationLen = 32

// Watcher provides asynchronous notifications on input changes
// The user should supply it pins to watch with AddPin and then wait for changes with Watch
// Alternately, users may receive directly from the Notification channel
// CN Watcher提供有关输入更改的异步通知，用户应向其提供引脚以使用AddPin进行监视，然后等待Watch的更改，或者，用户可以直接从Notification通道接收
type Watcher struct {
	pins         map[uintptr]Pin
	fds          fdHeap
	cmdChan      chan watcherCmd
	Notification chan WatcherNotification
}

// NewWatcher creates a new Watcher instance for asynchronous inputs
// CN 为异步输入创建一个新的Watcher实例
func NewWatcher() *Watcher {
	w := &Watcher{
		pins:         make(map[uintptr]Pin),
		fds:          fdHeap{},
		cmdChan:      make(chan watcherCmd, watcherCmdChanLen),
		Notification: make(chan WatcherNotification, notificationLen),
	}
	heap.Init(&w.fds)
	go w.watch()
	return w
}

func (w *Watcher) notify(fdset *syscall.FdSet) {
	for _, fd := range w.fds {
		if (fdset.Bits[fd/64] & (1 << (uint(fd) % 64))) != 0 {
			pin := w.pins[fd]
			val, err := pin.Read()
			if err != nil {
				if err == io.EOF {
					w.removeFd(fd)
					continue
				}
				fmt.Printf("failed to read pinfile, %s", err)
				os.Exit(1)
			}
			msg := WatcherNotification{
				Pin:   pin.Number,
				Value: val,
			}
			select {
			case w.Notification <- msg:
			default:
			}
		}
	}
}

func (w *Watcher) fdSelect() {
	timeval := &syscall.Timeval{
		Sec:  1,
		Usec: 0,
	}
	fdset := w.fds.FdSet()
	changed, err := doSelect(int(w.fds[0])+1, nil, nil, fdset, timeval)
	if err != nil {
		fmt.Printf("failed to call syscall.Select, %s", err)
		os.Exit(1)
	}
	if changed {
		w.notify(fdset)
	}
}

func (w *Watcher) addPin(p Pin) {
	fd := p.f.Fd()
	w.pins[fd] = p
	heap.Push(&w.fds, fd)
}

func (w *Watcher) removeFd(fd uintptr) {
	// heap operates on an array index, so search heap for fd
	// CN 堆对数组索引进行操作，因此在堆上搜索fd
	for index, v := range w.fds {
		if v == fd {
			heap.Remove(&w.fds, index)
			break
		}
	}
	pin := w.pins[fd]
	pin.f.Close()
	delete(w.pins, fd)
}

// removePin is only a wrapper around removeFd it finds fd given pin and then calls removeFd
// CN removePin只是removeFd的包装，它找到给定引脚的fd，然后调用removeFd
func (w *Watcher) removePin(p Pin) {
	// we don't index by pin, so go looking
	//CN 我们不按照pin索引，所以。。。
	for fd, pin := range w.pins {
		if pin.Number == p.Number {
			// found pin 找到pin
			w.removeFd(fd)
			return
		}
	}
}

func (w *Watcher) doCmd(cmd watcherCmd) (shouldContinue bool) {
	shouldContinue = true
	switch cmd.action {
	case watcherAdd:
		w.addPin(cmd.pin)
	case watcherRemove:
		w.removePin(cmd.pin)
	case watcherClose:
		shouldContinue = false
	}
	return shouldContinue
}

func (w *Watcher) recv() (shouldContinue bool) {
	for {
		select {
		case cmd := <-w.cmdChan:
			shouldContinue = w.doCmd(cmd)
			if !shouldContinue {
				return
			}
		default:
			shouldContinue = true
			return
		}
	}
}

func (w *Watcher) watch() {
	for {
		// first we do a syscall.select with timeout if we have any fds to check
		if len(w.fds) != 0 {
			w.fdSelect()
		} else {
			// so that we don't churn when the fdset is empty, sleep as if in select call
			time.Sleep(1 * time.Second)
		}
		if w.recv() == false {
			return
		}
	}
}

// AddPin adds a new pin to be watched for changes.
// The pin is configured with logic level "active high" and watched for both rising and falling edges.
// The pin provided should be the pin known by the kernel
// CN AddPin添加一个新的引脚，以监视其更改。
// CN 引脚配置为逻辑电平“高电平有效”，并注意上升沿和下降沿。
// CN 提供的引脚应该是内核已知的引脚
func (w *Watcher) AddPin(p uint) {
	w.AddPinWithEdgeAndLogic(p, EdgeBoth, ActiveHigh)
}

// AddPinWithEdgeAndLogic adds a new pin to be watched for changes.
// Edges can be configured to be either rising, falling, or both.
// Logic level can be active high or active low.
// The pin provided should be the pin known by the kernel.

// CN AddPinWithEdgeAndLogic添加一个新的引脚，以监视更改。
// CN 可以将边配置为上升，下降或同时上升和下降。 逻辑电平可以为高电平有效或低电平有效。
// CN 提供的引脚应该是内核已知的引脚。
func (w *Watcher) AddPinWithEdgeAndLogic(p uint, edge Edge, logicLevel LogicLevel) {
	pin := NewInput(p)
	setLogicLevel(pin, logicLevel)
	setEdgeTrigger(pin, edge)
	w.cmdChan <- watcherCmd{
		pin:    pin,
		action: watcherAdd,
	}
}

// RemovePin stops the watcher from watching the specified pin
// CN RemovePin阻止观察者观察指定的引脚
func (w *Watcher) RemovePin(p uint) {
	pin := Pin{
		Number: p,
	}
	w.cmdChan <- watcherCmd{
		pin:    pin,
		action: watcherRemove,
	}
}

// Watch blocks until one change occurs on one of the watched pins
// It returns the pin which changed and its new value
// Because the Watcher is not perfectly realtime it may miss very high frequency changes
// If that happens, it's possible to see consecutive changes with the same value
// Also, if the input is connected to a mechanical switch, the user of this library must deal with debouncing
// Users can either use Watch() or receive from Watcher.Notification directly

// CN 监视区块，直到其中一个被监视的引脚发生一次更改
// CN 它返回已更改的引脚及其新值，因为Watcher并非实时，因此可能会错过非常高的频率更改
// CN 如果发生这种情况，则可能会看到具有相同值的连续更改。此外，如果将输入连接到机械开关，则该库的用户必须处理反跳操作。
// CN 用户可以使用Watch（）或从Watcher接收。
func (w *Watcher) Watch() (p uint, v uint) {
	notification := <-w.Notification
	return notification.Pin, notification.Value
}

// Close stops the watcher and releases all resources
func (w *Watcher) Close() {
	w.cmdChan <- watcherCmd{
		pin:    Pin{},
		action: watcherClose,
	}
}
