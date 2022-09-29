package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
)

const layout = "2006-01-02 15:04:05"
const layoutMin = "2006-01-02 15:04"

var t = time.Now().Local()

// 获取客户端时间 并返回Info Gzs的值到客户端
type ClientData struct {
	Times string `json:"times"` // 客户端时间
	Gzs   string `json:"gzs"`   // 返回到客户端
	Info  string `json:"info"`  // 返回到客户端
}

// GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o server-arm64 serve.go
func main() {
	http.HandleFunc("/", home)
	//http.HandleFunc("/info", info)
	fmt.Println("serve @ localhost:6714")
	err := http.ListenAndServe(":6714", nil)
	if err != nil {
		log.Println(err)
	}
}
func home(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Printf("Method:%v\n", r.Method)
		resp(w)
	case "POST":
		fmt.Printf("Method:%v\n", r.Method)
		// parse json from client
		var clientData ClientData
		err := json.NewDecoder(r.Body).Decode(&clientData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		// clientData:  2022-09-29 11:26:25 len(19) 初始时间
		// clientData:  2022-09-30 11:26 len(16) 选择时间
		fmt.Printf("Get ClientData's times:  %v len(%d)\n", clientData.Times, len(clientData.Times))
		//返回到客户端
		var tx time.Time
		if len(clientData.Times) == 19 {
			tx, err = time.Parse(layout, clientData.Times)
			if err != nil {
				log.Println(err)
			}
		} else {
			tx, err = time.Parse(layoutMin, clientData.Times)
			if err != nil {
				log.Println(err)
			}
		}

		gzo := gz.NewTMGanZhi(tx.Year(), int(tx.Month()), tx.Day(), tx.Hour(), tx.Minute())
		info := gzo.Info()
		clientData.Gzs = info.GanZhiString
		clientData.Info = info.String()
		//
		json.NewEncoder(w).Encode(clientData)

	}
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
	//fmt.Println(string(b))
}

// func info(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("Method:%v\n", r.Method) //GET
// 	switch r.Method {
// 	case "GET":
// 		fmt.Println("get...", r.URL.Path)
// 		respInfo(w)
// 		// fmt.Fprintf(w, jsons, html.EscapeString(r.URL.Path))
// 	case "POST":
// 		fmt.Println("POST")
// 	}

// }
// func respInfo(w http.ResponseWriter) {
// 	gzo := gz.NewTMGanZhi(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
// 	info := gzo.Info()
// 	info.T = time.Date(1981, 1, 1, 1, 1, 1, 1, time.Local)
// 	b, err := json.Marshal(&info)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)

// 	} else {
// 		w.Header().Set("content-type", "application/json; charset=utf-8")
// 		w.Write(b)
// 	}
// 	fmt.Println("info---> ", string(b))
// }
