package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init() {
	//	_ = os.Setenv("FYNE_FONT", "font/SourceHanSerifCN-SemiBold.ttf")
	_ = os.Setenv("FYNE_FONT", "/system/fonts/SourceHanSerifCN-SemiBold.ttf") //apk字体
}
func main() {
	a := app.New()
	w := a.NewWindow("演禽")
	a.Settings().SetTheme(theme.LightTheme())

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

	label := widget.NewLabel("")
	label3 := widget.NewLabel("")
	label2 := widget.NewLabel("")
	label1 := widget.NewLabel("")

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
		year, err := strconv.Atoi(ys)
		if err != nil {
			fmt.Println(err)
		}
		month := selectm.SelectedIndex() + 1
		day := selectd.SelectedIndex() + 1
		hour := selecth.SelectedIndex()
		//显示到主页面
		ymdh := fmt.Sprintf("%d年-%d月-%d日-%d时", year, month, day, hour)
		label.SetText(ymdh)
		yq := get(year, month, day, hour)

		label3.SetText(yq.GetGanZhi() + "\n" + yq.GetQiKe() + "\n" + yq.GetJiangTou() + "\n" + yq.GetFanQinHuoYao())
		xjbf := getXjbf(yq.Ygz, yq.Mgz, yq.Dgz, yq.Hgz)
		label2.SetText(xjbf.JianChu())
		label1.SetText(yq.GetSimpleJX())
	})
	////img
	//img := canvas.NewImageFromFile("img/震邪符咒.png")
	//img.FillMode = canvas.ImageFillContain
	//c := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 150)), img)

	//
	hbox := container.New(layout.NewHBoxLayout(), ent, selectm, selectd, selecth, btnOK)

	//
	vbox := container.New(layout.NewVBoxLayout(), label, label3, label2, label1)

	w.SetContent(container.New(layout.NewVBoxLayout(), hbox, vbox))
	w.Resize(fyne.Size{600, 600})
	w.ShowAndRun()

}
