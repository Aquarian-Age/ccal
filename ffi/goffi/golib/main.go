package main

//import "github.com/Aquarian-Age/sjqm/v3"
import "time"
import "fmt"

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// go build -o libgets.so -buildmode=c-shared main.go
func main() {}

//export add
func add(a, b int) int {
	return a + b
}

//export gets
func gets() *C.char {
	days := day_week()

	s := "test  go string to c *char after day_week()" + days
	defer C.free(unsafe.Pointer(C.CString(s)))
	return C.CString(s)
}
func day_week() string {
	fmt.Println("-------------------------------")
	//sjqm.FindRiNei(time.Now().Local(), "", "", 2)
	return "这是日内查找....." + time.Now().Local().String()
}
