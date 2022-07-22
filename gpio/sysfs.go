package gpio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type direction uint

const (
	inDirection direction = iota
	outDirection
)

type Edge uint

const (
	EdgeNone Edge = iota
	EdgeRising
	EdgeFalling
	EdgeBoth
)

type LogicLevel uint

const (
	ActiveHigh LogicLevel = iota
	ActiveLow
)

type Value uint

const (
	Inactive Value = 0
	Active   Value = 1
)

//????
func exportGPIO(p Pin) {
	export, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("exportGPIO-->failed to open gpio export file for writing\n")
		os.Exit(1)
	}
	defer export.Close()
	export.Write([]byte(strconv.Itoa(int(p.Number))))
}

//????
func unexportGPIO(p Pin) {
	export, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("unexportGPIO-->failed to open gpio unexport file for writing\n")
		os.Exit(1)
	}
	defer export.Close()
	export.Write([]byte(strconv.Itoa(int(p.Number))))
}

//????????
func setDirection(p Pin, d direction, initialValue uint) {
	dir, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpiochip%d/direction", p.Number), os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("setDirection-->failed to open gpiochip %d direction file for writing\n", p.Number)
		os.Exit(1)
	}
	defer dir.Close()

	switch {
	case d == inDirection:
		dir.Write([]byte("in"))
	case d == outDirection && initialValue == 0:
		dir.Write([]byte("low"))
	case d == outDirection && initialValue == 1:
		dir.Write([]byte("high"))
	default:
		panic(fmt.Sprintf("setDirection-->setDirection called with invalid direction or initialValue, %d, %d", d, initialValue))
	}
}

func setEdgeTrigger(p Pin, e Edge) {
	edge, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpiochip%d/edge", p.Number), os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("setEdgeTrigger-->failed to open gpiochip %d edge file for writing\n", p.Number)
		os.Exit(1)
	}
	defer edge.Close()

	switch e {
	case EdgeNone:
		edge.Write([]byte("none"))
	case EdgeRising:
		edge.Write([]byte("rising"))
	case EdgeFalling:
		edge.Write([]byte("falling"))
	case EdgeBoth:
		edge.Write([]byte("both"))
	default:
		panic(fmt.Sprintf("setEdgeTrigger-->setEdgeTrigger called with invalid edge %d", e))
	}
}

// ?????????????????????????????????????????????????????
func setLogicLevel(p Pin, l LogicLevel) error {
	level, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpiochip%d/active_low", p.Number), os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer level.Close()

	switch l {
	case ActiveHigh:
		level.Write([]byte("0"))
	case ActiveLow:
		level.Write([]byte("1"))
	default:
		return errors.New("setLogicLevel-->Invalid logic level setting.")
	}
	return nil
}

//????
func openPin(p Pin, write bool) Pin {
	flags := os.O_RDONLY
	if write {
		flags = os.O_RDWR
	}
	f, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpiochip%d/label", p.Number), flags, 0600) //rk??label base
	if err != nil {
		fmt.Printf("setEdgeTrigger-->failed to open gpiochip %d value file for reading\n", p.Number)
		os.Exit(1)
	}
	p.f = f
	return p
}

//???? ??0??1 ??
func readPin(p Pin) (val uint, err error) {
	file := p.f
	file.Seek(0, 0)
	buf := make([]byte, 1)
	_, err = file.Read(buf)
	if err != nil {
		return 0, err
	}
	c := buf[0]
	switch c {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	default:
		return 0, fmt.Errorf("readPin-->read inconsistent value in pinfile, %c", c)
	}
}

//??? ????
func writePin(p Pin, v uint) error {
	var buf []byte
	switch v {
	case 0:
		buf = []byte{'0'}
	case 1:
		buf = []byte{'1'}
	default:
		return fmt.Errorf("writePin-->invalid output value %d", v)
	}
	_, err := p.f.Write(buf)
	return err
}
