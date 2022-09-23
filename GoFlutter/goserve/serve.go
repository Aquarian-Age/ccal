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

type ClientData struct {
	Times string `json:"times"`
}

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
	fmt.Printf("Method:%v\n", r.Method) //GET
	switch r.Method {
	case "GET":
		fmt.Println("get...", r.URL.Path)
		resp(w)
	case "POST":
		fmt.Println("POST")
		// parse json from client
		var clientData ClientData
		err := json.NewDecoder(r.Body).Decode(&clientData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		fmt.Printf("clientData:  %v\n", clientData.Times)
		json.NewEncoder(w).Encode(clientData) //返回到客户端
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
	fmt.Println(string(b))
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method:%v\n", r.Method) //GET
	switch r.Method {
	case "GET":
		fmt.Println("get...", r.URL.Path)
		respInfo(w)
		// fmt.Fprintf(w, jsons, html.EscapeString(r.URL.Path))
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

// get... /
//{"t":"2022-09-21T23:09:00+08:00","ygz":"壬寅","mgz":"己酉","dgz":"丁丑","hgz":"庚子","moon":"八月廿六","solar":"2022-09-21 23:09:00","lu":"禄位: 亥 - 午 - 午 - 庚","gan_zhi":"干支: 壬寅 己酉 丁丑 庚子","xun_kong":"旬空: 辰巳 寅卯 申酉 辰巳","na_yin":"纳音: 金箔金-大鄢土-涧下水-壁上土","zhong_qi_string":"中气: 处暑(2022-08-23)","yue_jiang_string":"月将: 巳(太乙)","jian_chu":"建除: 定","ri_qin":"日禽: 轸水蚓","huang_hei":"黄黑: 勾陈","huang_hei_h":"黄黑H: 天刑","jue_li_ri":"","ji_tan_bing":""}
//Method:GET
//get... /info
//info--->  {"t":"1981-01-01T01:01:01.000000001+08:00","ygz":"壬寅","mgz":"己酉","dgz":"丁丑","hgz":"庚子","moon":"八月廿六","solar":"2022-09-21 23:09:00","lu":"禄位: 亥 - 午 - 午 - 庚","gan_zhi":"干支: 壬寅 己酉 丁丑 庚子","xun_kong":"旬空: 辰巳 寅卯 申酉 辰巳","na_yin":"纳音: 金箔金-大鄢土-涧下水-壁上土","zhong_qi_string":"中气: 处暑(2022-08-23)","yue_jiang_string":"月将: 巳(太乙)","jian_chu":"建除: 定","ri_qin":"日禽: 轸水蚓","huang_hei":"黄黑: 勾陈","huang_hei_h":"黄黑H: 天刑","jue_li_ri":"","ji_tan_bing":""}
