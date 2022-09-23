package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
)

var t = time.Now().Local()

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/info", info)
	fmt.Println("serve @ localhost:6714")
	err := http.ListenAndServe(":6714", nil)
	if err != nil {
		log.Println(err)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("get...", r.URL.Path)
		resp(w)
	case "POST":
		fmt.Println("POST")
	}

}
func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("get...", r.URL.Path)
		respInfo(w)
	case "POST":
		fmt.Println("POST")
	}

}
func respInfo(w http.ResponseWriter) {
	gzo := gz.NewTMGanZhi(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
	info := gzo.Info()
	info.T = time.Date(1981, 1, 1, 1, 1, 1, 1, time.Local)
	b, err := json.Marshal(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.Write(b)
	}
	fmt.Println("info---> ", string(b))
}
func resp(w http.ResponseWriter) {
	gzo := gz.NewTMGanZhi(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
	info := gzo.Info()
	b, err := json.Marshal(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.Write(b)
	}
	fmt.Println(string(b))
}
