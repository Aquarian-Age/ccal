package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

//默认启用9090端口
func main() {
	flag.Usage = func() {
		fmt.Println("usage ccal-web -l port")
		flag.PrintDefaults()
	}
	p := flag.Int("l", 80, "")
	flag.Parse()

	fmt.Println("https://github.com/Aquarian-Age/ccal-gui/tree/master/web 下载webUI页面")

	http.HandleFunc("/", home)
	http.HandleFunc("/jz60", selectlist)
	//http.HandleFunc("/yiji", yiJi)

	port := fmt.Sprintf(":%d", *p)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
