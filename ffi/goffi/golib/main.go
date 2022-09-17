package main

//import "github.com/Aquarian-Age/sjqm/v3"
import (
	"encoding/json"
	"log"
	"time"
	"unsafe"
)

// #include <stdio.h>
// #include <stdlib.h>
import "C"

// go build -o libgets.so -buildmode=c-shared main.go
func main() {
}

//export add
func add(a, b int) int {
	return a + b
}

//export gets
func gets() *C.char {
	days := day_week()

	//s := "test  go string to c *char " + days
	s := days
	defer C.free(unsafe.Pointer(C.CString(s)))
	return C.CString(s)
}
func day_week() string {
	group := &Group{
		Id:   "group id",
		Name: "group name",
		Time: time.Now().Local().String(),
	}
	b, err := json.Marshal(group)
	if err != nil {
		log.Println(err)
	}

	return string(b)
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}
