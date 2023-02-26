package main

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()
	p := gpioreg.ByName("11")
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.Low; ; l = !l {
		p.Out(l)
		<-t.C
	}
}
