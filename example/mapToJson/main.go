package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("home.html")
		if err != nil {
			fmt.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			return
		}

		body := newBody()
		b, _ := json.Marshal(body)
		err = json.NewEncoder(w).Encode(string(b))
		if err != nil {
			return
		}

	}
}

type Body struct {
	Map map[string]string `json:"map"`
}

func newBody() *Body {
	arr := []string{"天机星", "太阳星", "武曲星", "天同星", "廉贞星",
		"天府星", "太阴星", "贪狼星", "巨门星", "天相星", "天梁星", "七杀星", "破军星",
		"紫微星"}
	zhi := []string{"子", "丑", "寅", "寅", "卯", "辰", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	var xmap = make(map[string]string)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(zhi); j++ {
			if i == j {
				xmap[arr[j]] = zhi[i]
				break
			}
		}
	}
	return &Body{
		Map: xmap,
	}
}
