package main

import (
	"fmt"
	"log"

	"github.com/fraburnham/gpio/sysfs"
)

func main() {
	io, err := sysfs.NewInput(9)
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.ReadValue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("gpio value:%d\n", n)

}
