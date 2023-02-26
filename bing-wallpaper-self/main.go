package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var p = flag.String("p", "/home/user/Picture/", "")

func main() {

	flag.Parse()

	r, err := Get(0, "zh-CN", "1920")
	if err != nil {
		log.Println(err)
	}

	resp, err := http.Get(r.URL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if *p != "" {
		name := fmt.Sprintf("%s%s.jpg", *p, r.EndDate)
		fmt.Println(name)
		ioutil.WriteFile(name,b, 0666)
	} else {
		ioutil.WriteFile(r.EndDate+".jpg", b, 0666)
	}
}
