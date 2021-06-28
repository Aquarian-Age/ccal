package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/starainrt/astro/basic"
	"liangzi.local/qx"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"liangzi.local/qx/pkg/cal"
)

func main() {
	rect := sciter.Rect{Left: 0, Top: 0, Right: 1200, Bottom: 750}
	w, err := window.New(
		sciter.SW_TITLEBAR|
			sciter.SW_RESIZEABLE|
			sciter.SW_CONTROLS|
			sciter.SW_MAIN|
			sciter.SW_ENABLE_DEBUG,
		&rect)
	if err != nil {
		log.Fatal("window:", err)
	}
	//w.LoadFile("cal.html")
	w.LoadHtml(html, "")

	setWindowHandler(w)
	w.Show()
	w.Run()
}

func setWindowHandler(w *window.Window) {
	w.DefineFunction("ymdinfo", ymdinfo)
	w.DefineFunction("yuelibiao", yuelibiao)
}

//纪年信息
type JN struct {
	Sjn string `json:"sjn"` //阳历纪年
	Ljn string `json:"ljn"` //阴历纪年
	Gjn string `json:"gjn"` //干支纪年
	Lmb string `json:"lmb"` //演禽
	Ny  string `json:"ny"`  //纳因
}
type YueLiBiao struct {
	SDayArr     []string `json:"sdays"`
	LDayArr     []string `json:"ldays"`
	GZhiArr     []string `json:"gzhis"`
	JianChuArr  []string `json:"jian_chu_arr"`
	HuangHeiArr []string `json:"huang_hei_arr"`
	StarArr     []string `json:"star_arr"`
	S           string   `json:"s"`
}

//纪年信息
func ymdinfo(args ...*sciter.Value) *sciter.Value {
	ly, lm, ld, lh := args[0].String(), args[1].String(), args[2].String(), args[3].String()
	y, m, d, h := ConvStoInt(ly, lm, ld, lh)
	obj := qx.NewYanQin(y, m, d, h)
	ygz := obj.Ygz
	mgz := obj.Mgz
	dgz := obj.Dgz
	hgz := obj.Hgz

	_, _, _, moons := basic.GetLunar(y, m, d)
	ygzny := gz.GetNaYin(ygz)
	mgzny := gz.GetNaYin(mgz)
	dgzny := gz.GetNaYin(dgz)
	hgzny := gz.GetNaYin(hgz)

	jns := fmt.Sprintf("阳历:%d年-%d月-%d日:%d时 周%d", obj.Year, obj.Month, obj.Day, obj.Hour, obj.Week)
	jnl := fmt.Sprintf("阴历: %s", moons)
	jng := fmt.Sprintf("干支:%s年-%s月-%s日-%s时", ygz, mgz, dgz, hgz)
	ny := fmt.Sprintf("纳音: %s %s %s %s", ygzny, mgzny, dgzny, hgzny)
	jnmb := fmt.Sprintf("%s-%s-%s-%s", obj.NianQin, obj.YueQin, obj.RiQin, obj.ShiQin)

	jn := JN{
		Sjn: jns,
		Ljn: jnl,
		Gjn: jng,
		Lmb: jnmb,
		Ny:  ny,
	}
	jnjson, err := json.Marshal(jn)
	if err != nil {
		log.Fatal(err)
	}
	return sciter.NewValue(string(jnjson))
}
func ConvStoInt(ys, ms, ds, hs string) (int, int, int, int) {
	y, err := strconv.Atoi(ys)
	if err != nil {
		log.Fatal("年份時間解析:", err)
	}

	m, err := strconv.Atoi(ms)
	if err != nil {
		log.Fatal("月份時間解析:", err)
	}
	d, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal("日期時間解析:", err)
	}
	h, err := strconv.Atoi(hs)
	if err != nil {
		log.Fatal("時辰解析:", err)
	}

	return y, m, d, h
}

//月历
func yuelibiao(args ...*sciter.Value) *sciter.Value {
	ly, lm, ld, lh := args[0].String(), args[1].String(), args[2].String(), args[3].String()
	y, m, d, h := ConvStoInt(ly, lm, ld, lh)
	lunarArr := cal.NewLunarArr(y, m, h)
	sdays, ldays, gzdays, starArr := lunarArr.SolarArr, lunarArr.MoonArr, lunarArr.GzArr, lunarArr.StarArr
	jianChuArr := lunarArr.RjcArr
	huangHeiArr := lunarArr.RiHuangHeiArr

	t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	st := fmt.Sprintf("%d年%d月%d日%d点 周%d", y, m, d, h, int(t.Weekday()))
	ylbiao := YueLiBiao{
		SDayArr:     sdays,
		LDayArr:     ldays,
		GZhiArr:     gzdays,
		JianChuArr:  jianChuArr,
		HuangHeiArr: huangHeiArr,
		StarArr:     starArr,
		S:           st,
	}
	//fmt.Printf("%v\n", ylbiao)
	ylb, err := json.Marshal(ylbiao)
	if err != nil {
		log.Println("ylb:", err)
	}
	return sciter.NewValue(string(ylb))
}
