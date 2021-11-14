package main

import (
	"fmt"
	"liangzi.local/qx/pkg/bmqd"
	"liangzi.local/qx/pkg/gz"
	"liangzi.local/qx/pkg/x"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("演禽")                    //标题
	a.Settings().SetTheme(&Biu{})

	t := time.Now().Local()
	ys := fmt.Sprintf("%d", t.Year())
	ent := widget.NewEntry()
	ent.SetText(ys)

	//select options
	months := func() []string {
		var s []string
		for i := 1; i <= 12; i++ {
			is := fmt.Sprintf("%d月", i)
			s = append(s, is)
		}
		return s
	}
	days := func() []string {
		var s []string
		for i := 1; i <= 31; i++ {
			is := fmt.Sprintf("%d日", i)
			s = append(s, is)
		}
		return s
	}

	hours := func() []string {
		var s []string
		for i := 0; i <= 23; i++ {
			is := fmt.Sprintf("%d时", i)
			s = append(s, is)
		}
		return s
	}
	selectm := widget.NewSelect(months(), func(s string) {
	})
	selectm.PlaceHolder = "month"
	selectd := widget.NewSelect(days(), func(s string) {
	})
	selectd.PlaceHolder = "day"
	selecth := widget.NewSelect(hours(), func(s string) {
	})
	selecth.PlaceHolder = "hour"

	labelSolar := widget.NewLabel("")
	labelXjbfs := widget.NewLabel("")
	//labelSolar.Wrapping = fyne.TextWrapWord //长文本自动换行
	//labelXjbfs.Wrapping = fyne.TextWrapWord //长文本自动换行
	labelSolar.Wrapping=fyne.TextWrapBreak
	labelXjbfs.Wrapping=fyne.TextWrapBreak

	btnOK := widget.NewButton("OK", func() {
		ys := ent.Text
		if selectm.SelectedIndex() == -1 {
			selectm.SetSelectedIndex(int(t.Month()) - 1)
		}
		if selectd.SelectedIndex() == -1 {
			selectd.SetSelectedIndex(t.Day() - 1)
		}
		if selecth.SelectedIndex() == -1 {
			selecth.SetSelectedIndex(t.Hour())
		}
		//这里获取下拉框选中的值
		year, err := strconv.Atoi(ys)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		day := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		t = time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Local)

		//显示到主页面
		solarS := fmt.Sprintf("%d年-%d月-%d日 ", year, month, day )

		gzo:=gz.NewGanZhi(year,month,day,hour)
		naYinS :=fmt.Sprintf("%s\n", gzo.GetNaYin())
		lunar, moon := gzo.GetLunar()
		lunarS := fmt.Sprintf("%s %s\n", lunar, moon)
		jieQiS := gzo.JieQi()
		ganZhiS := fmt.Sprintf("%s %s %s %s\n", gzo.YGZ, gzo.MGZ, gzo.DGZ, gzo.HGZ )
		yjzhi, yjName, _ := gzo.YueJiang()
		yueJiangS := fmt.Sprintf("节气:%s 月将:%s(%s)\n\n",jieQiS, yjzhi, yjName)
		//
		jianChu := gzo.RiJianChu()
		jianChuInfo := x.JianChuSelf(jianChu)
		jianChuS := fmt.Sprintf("日建除:%s %s\n", jianChu,jianChuInfo)

		huangHei := gzo.RiHuangHei()
		hhInfo := gz.HuangHeiJX(huangHei) //鼠标悬停提示
		huangHeiS:=fmt.Sprintf("日黄黑:%s %s\n",huangHei,hhInfo)

		riQin := gzo.RiQin(int(t.Weekday()))
		starOwner := bmqd.GetOwner(riQin)
		shuaiWangS := bmqd.ShuaiWang(riQin)
		starSelf := x.StarSelf(riQin)
		starHint := fmt.Sprintf("%s %s\n%s\n",starOwner ,shuaiWangS ,starSelf)
		riQinS := fmt.Sprintf("日禽:%s %s\n", riQin,starHint)

		huangHeiH := gzo.ShiHuangHei()
		hourHuangHeiS := fmt.Sprintf("时黄黑:%s\n", huangHeiH)
		labelSolar.SetText(solarS+lunarS+ganZhiS+naYinS+
			yueJiangS+jianChuS+huangHeiS +riQinS+
		hourHuangHeiS)

		jiaZiZhengXiangS := bmqd.JiaZiZhengXiang(gzo.DGZ)
		labelXjbfs.SetText(jiaZiZhengXiangS)
	})
	//基础部分
	hbox := container.New(layout.NewHBoxLayout(), ent, selectm, selectd, selecth, btnOK)
	vbox := container.New(layout.NewVBoxLayout(), labelSolar, labelXjbfs)
	basic := container.New(layout.NewVBoxLayout(), hbox, vbox)
	dw := container.New(layout.NewHBoxLayout(), basic)
	w.SetContent(dw)
	w.Resize(fyne.Size{Width: 380, Height: 420})
	w.ShowAndRun()
}
