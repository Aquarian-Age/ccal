package main

import (
	"fmt"

	"github.com/brian-armstrong/gpio"
)

func watch() {

	//读取指定GPIO数字的状态
	P := gpio.NewOutput(252, true)
	fmt.Printf("GPIO数字 值:%v\n", P)

	//指定GPIO数字的输入状态
	rn, err := P.Read()
	if err != nil {
		fmt.Printf("电平输入状态:%v\n", err)
	}
	fmt.Printf("电平输入数值为:%d\n", rn)

	///var i uint
	//for i = 0; i < 100; i++ {
	//pin := gpio.NewInput(6)
	//fmt.Printf("查看的pin:%v\n", pin)
	// watcher := gpio.NewWatcher()
	// fmt.Printf("watcher:%v\n", watcher)
	// fmt.Println("--------- andd 7 2 -------------")
	// //watcher.AddPin(7)
	// watcher.AddPin(216)
	// defer watcher.Close()

	// go func() {
	// 	for {
	// 		pin, value := watcher.Watch()
	// 		fmt.Printf("read %d from gpio %d\n", value, pin)
	// 	}
	// }()
	//}
}

func main() {
	watch()
}
