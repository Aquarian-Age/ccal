package main

import (
	"flag"
	"fmt"
	"liangzi.local/qx/web"
)

//默认8111 指定端口 yq -l 9999
func main() {
	flag.Usage = func() {
		fmt.Println("yq -l port")
		flag.PrintDefaults()
	}
	p := flag.Int("l", 8111, "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	fmt.Println("Default Server Port:", 8111)
	web.Main(port)
}
