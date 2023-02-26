package gpio

import (
	"errors"
	"os"
	"time"
)

// Pin represents a single pin, which can be used either for reading or writing
//CN: Pin代表单个引脚，可用于读取或写入
type Pin struct {
	Number    uint
	direction direction //方向
	f         *os.File
}

// NewInput opens the given pin number for reading. The number provided should be the pin number known by the kernel
//CN: 打开给定的引脚号以进行读取。提供的数字应该是内核知道的引脚号
func NewInput(p uint) Pin {
	pin := Pin{
		Number: p,
	}
	exportGPIO(pin)
	time.Sleep(10 * time.Millisecond)
	pin.direction = inDirection
	setDirection(pin, inDirection, 0)
	pin = openPin(pin, false)
	return pin
}

// NewOutput opens the given pin number for writing. The number provided should be the pin number known by the kernel
// NewOutput also needs to know whether the pin should be initialized high (true) or low (false)
// CN NewOutput打开给定的引脚号进行写入。 提供的数字应该是内核知道的引脚号
// CN NewOutput还需要知道应将引脚初始化为高（true）还是低（false）
func NewOutput(p uint, initHigh bool) Pin {
	pin := Pin{
		Number: p,
	}
	exportGPIO(pin)
	time.Sleep(10 * time.Millisecond)
	initVal := uint(0)
	if initHigh {
		initVal = uint(1)
	}
	pin.direction = outDirection
	setDirection(pin, outDirection, initVal)
	pin = openPin(pin, true)
	return pin
}

// Close releases the resources related to Pin. This doen't unexport Pin, use Cleanup() instead
//CN 关闭释放与Pin相关的资源。这不会取消导出Pin，请使用Cleanup（）代替
func (p Pin) Close() {
	if p.f != nil {
		p.f.Close()
		p.f = nil
	}
}

// Cleanup close Pin and unexport it
// CN 清理关闭Pin并取消导出
func (p Pin) Cleanup() {
	p.Close()
	unexportGPIO(p)
}

// Read returns the value read at the pin as reported by the kernel. This should only be used for input pins
// CN 返回内核报告的从引脚读取的值。仅应用于输入引脚
func (p Pin) Read() (value uint, err error) {
	if p.direction != inDirection {
		return 0, errors.New("Read()-->pin is not configured for input")
	}
	return readPin(p)
}

// SetLogicLevel sets the logic level for the Pin. This can be either "active high" or "active low"
// CN 设置引脚的逻辑电平。可以是“高电平有效”或“低电平有效”
func (p Pin) SetLogicLevel(logicLevel LogicLevel) error {
	return setLogicLevel(p, logicLevel)
}

// High sets the value of an output pin to logic high
// CN 高电平将输出引脚的值设置为逻辑高电平
func (p Pin) High() error {
	if p.direction != outDirection {
		return errors.New("High()-->pin is not configured for output")
	}
	return writePin(p, 1)
}

// Low sets the value of an output pin to logic low
// CN 低电平将输出引脚的值设置为逻辑低电平
func (p Pin) Low() error {
	if p.direction != outDirection {
		return errors.New("Low()-->pin is not configured for output")
	}
	return writePin(p, 0)
}
