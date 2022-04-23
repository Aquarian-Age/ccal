package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"liangzi.local/qx/pkg/pub"
	"liangzi.local/qx/pkg/yq"
	"strconv"
	"time"
)

type UI struct {
}

var (
	list = []string{"男子", "妇女", "小儿", "女子",
		"牛", "马", "羊", "猪", "驴", "骡", "鸡", "狗",
		"衣服", "金银"}
	resultS string
	t       = time.Now().Local()
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&Biu{})
	w := a.NewWindow("二十八宿走失诀")

	labels := widget.NewLabel("")
	ys := strconv.Itoa(t.Year())
	ent := widget.NewEntry()
	ent.SetText(ys)

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
	//
	selectm := widget.NewSelect(months(), func(s string) {
	})
	selectm.PlaceHolder = "月"
	selectd := widget.NewSelect(days(), func(s string) {
	})
	selectd.PlaceHolder = "日"
	selecth := widget.NewSelect(hours(), func(s string) {
	})
	selecth.PlaceHolder = "时"
	selectItem := widget.NewSelect(list, func(s string) {
	})
	selectItem.PlaceHolder = "测事"
	self := selectItem.Selected

	btnOK := widget.NewButton("一侧", func() {
		yearS := ent.Text
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
		year, err := strconv.Atoi(yearS)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		dayd := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		n := selectItem.SelectedIndex()
		self = list[n]
		resultS = btnok(year, month, dayd, hour, self)
		labels.SetText(resultS)

	})
	btnOK1 := widget.NewButton("再测", func() {
		yearS := ent.Text
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
		year, err := strconv.Atoi(yearS)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		dayd := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		n := selectItem.SelectedIndex()
		self = list[n]
		resultS = btnok1(year, month, dayd, hour, self)
		labels.SetText(resultS)
	})
	btnOK2 := widget.NewButton("三测", func() {
		yearS := ent.Text
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
		year, err := strconv.Atoi(yearS)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		dayd := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		n := selectItem.SelectedIndex()
		self = list[n]
		resultS = btnok2(year, month, dayd, hour, self)
		labels.SetText(resultS)
	})
	hbox := container.New(layout.NewHBoxLayout(), ent, selectm, selectd, selecth, selectItem)
	hboxBtns := container.New(layout.NewHBoxLayout(), btnOK, btnOK1, btnOK2)
	vbox := container.New(layout.NewHBoxLayout(), labels)

	showa := container.New(layout.NewVBoxLayout(), hbox, hboxBtns, vbox)
	w.SetContent(showa)
	w.Resize(fyne.Size{Width: 360, Height: 640})
	w.ShowAndRun()
}

//首测
func btnok(y, m, d, h int, self string) string {
	t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wd := int(t.Weekday())
	wds := pub.WeekName(wd)
	dgz, _ := yq.GanZhiD(y, int(m), int(d))
	riQin := yq.GetRiQin(wd, dgz)
	hz := pub.ConvH24ToH12S(t.Hour())
	res := yq.ZouShiJue(riQin, hz, self)
	res = fmt.Sprintf("失物方向: %s\n", res)
	s := fmt.Sprintf("%d年%d月%d日%d点 %s\n", y, m, d, h, wds)
	s1 := fmt.Sprintf("日宿:%s\n%s时起课\n测:%s\n", riQin, hz, self)

	res = s + s1 + res
	kys := yq.ZouShiKeYing(self, hz)
	res += kys

	return res
}

//再测
func btnok1(y, m, d, h int, self string) string {
	t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wd := int(t.Weekday())
	wds := pub.WeekName(wd)
	dgz, _ := yq.GanZhiD(y, int(m), int(d))
	riQin := yq.GetRiQin(wd, dgz)
	hz := pub.ConvH24ToH12S(t.Hour())
	res := yq.ZouShiJue2(riQin, hz, self)
	res = fmt.Sprintf("失物方向: %s\n", res)
	s := fmt.Sprintf("%d年%d月%d日%d点 %s\n", y, m, d, h, wds)
	s1 := fmt.Sprintf("日宿:%s\n%s时起课\n测:%s\n", riQin, hz, self)

	res = s + s1 + res
	kys := yq.ZouShiKeYing(self, hz)
	res += kys
	return res
}

//三测
func btnok2(y, m, d, h int, self string) string {
	t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wd := int(t.Weekday())
	wds := pub.WeekName(wd)
	dgz, _ := yq.GanZhiD(y, int(m), int(d))
	riQin := yq.GetRiQin(wd, dgz)
	hz := pub.ConvH24ToH12S(t.Hour())
	res := yq.ZouShiJue3(riQin, hz, self)
	res = fmt.Sprintf("失物方向: %s\n", res)
	s := fmt.Sprintf("%d年%d月%d日%d点 %s\n", y, m, d, h, wds)
	s1 := fmt.Sprintf("日宿:%s\n%s时起课\n测:%s\n", riQin, hz, self)

	res = s + s1 + res
	kys := yq.ZouShiKeYing(self, hz)
	res += kys
	return res
}
