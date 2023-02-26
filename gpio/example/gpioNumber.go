//https://periph.io/
package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO6")
	if p == nil {
		log.Fatal("Failed to find GPIO6")
	}

	// A pin can be read, independent of its state; it doesn't matter if it is
	// set as input or output.
	fmt.Printf("%s is %s\n", p, p.Read())
}
