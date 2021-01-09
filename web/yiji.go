package main

/* //吉凶 宜忌
func yiJi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("yiji.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		yinfo := Yearinfo()

		json.NewEncoder(w).Encode(yinfo)
	}
} */

/* //年煞 太岁出游 金神 五鬼 破败五鬼
type YearJX struct {
	TaiSui                       string
	JinShen1, JinShen2, JinShen3 string
}

func Yearinfo() *YearJX {
	yjx := new(YearJX)
	taisui := []string{"太岁", "金神", "五鬼", "破败五鬼"}
	info := jx.NewJinShenWG(taisui)
	ts1,ts2,ts3 := info.TaiSui()
	ts := ts1+ts2+ts3
	j1, j2, j3 := info.JinShen()

	yjx = &YearJX{
		TaiSui:   ts,
		JinShen1: j1,
		JinShen2: j2,
		JinShen3: j3,
	}
	return yjx

} */
