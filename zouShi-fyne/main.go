package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"liangzi.local/qx/pkg/pub"
	"liangzi.local/qx/pkg/yq"
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

	months := []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
	days := []string{"1日", "2日", "3日", "4日", "5日", "6日", "7日", "8日", "9日", "10日", "11日", "12日", "13日", "14日", "15日", "16日", "17日", "18日", "19日", "20日", "21日", "22日", "23日", "24日", "25日", "26日", "27日", "28日", "29日", "30日", "31日"}
	hours := []string{"0点", "1点", "2点", "3点", "4点", "5点", "6点", "7点", "8点", "9点", "10点", "11点", "12点", "13点", "14点", "15点", "16点", "17点", "18点", "19点", "20点", "21点", "22点", "23点"}

	selectm := widget.NewSelect(months, func(s string) {})
	selectm.PlaceHolder = "月"
	selectd := widget.NewSelect(days, func(s string) {})
	selectd.PlaceHolder = "日"
	selecth := widget.NewSelect(hours, func(s string) {})
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
	//hbox := container.New(layout.NewGridWrapLayout(fyne.Size{60, 36}), ent, selectm, selectd, selecth, selectItem)
	gridSelect := container.New(layout.NewGridLayout(5), ent, selectm, selectd, selecth, selectItem)
	//hboxBtns := container.New(layout.NewGridWrapLayout(fyne.Size{60, 30}), btnOK, btnOK1, btnOK2)
	gridBtn := container.New(layout.NewGridLayout(3), btnOK, btnOK1, btnOK2)
	gridLabel := container.New(layout.NewGridLayout(1), labels)
	showa := container.New(layout.NewVBoxLayout(), gridSelect, gridBtn, gridLabel)
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
