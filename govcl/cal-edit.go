package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/starainrt/astro/basic"

	// 如果你使用自定义的syso文件则不要引用此包
	"time"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	label *vcl.TLabel
	year  *vcl.TEdit
	Btn1  *vcl.TButton
	Btn2  *vcl.TButton
	//
	edit, edit2, edit3, edit4, edit5, edit6, edit7         *vcl.TEdit
	edit8, edit9, edit10, edit11, edit12, edit13, edit14   *vcl.TEdit
	edit15, edit16, edit17, edit18, edit19, edit20, edit21 *vcl.TEdit
	edit22, edit23, edit24, edit25, edit26, edit27, edit28 *vcl.TEdit
	edit29, edit30, edit31                                 *vcl.TEdit
}

type TAboutForm struct {
	*vcl.TForm
	Btn1  *vcl.TButton
	edit  *vcl.TEdit
	label *vcl.TLabel
}

var (
	mainForm  *TMainForm
	aboutForm *TAboutForm
)

func main() {
	vcl.RunApp(&mainForm, &aboutForm)
}
func info(y, m, d, h int) string {
	solar := fmt.Sprintf("%d月%d日", m, d) //阳历
	//
	t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	wi := int(t.Weekday())
	warr := []string{"日", "一", "二", "三", "四", "五", "六"}
	var ws string
	for i := 0; i < len(warr); i++ {
		if wi == i {
			ws = "周" + warr[i]
			break
		}
	}
	//
	_, _, _, moon := basic.GetLunar(y, m, d) //阴历
	mgz := gz.GetMonthGZ(y, m, d, h)         //月干支
	dgz := gz.GetDayGZ(y, m, d)              //日干支
	jianChu := gz.GetRiJianChu(mgz, dgz)     //日建除
	huangHei := gz.GetRiHuangHei(mgz, dgz)   //日黄黑

	wd := int(t.Weekday())
	riQin := gz.GetRiQin(wd, dgz)

	s := solar + ws + " " + moon + " " + dgz + " " + jianChu + " " + huangHei + " " + riQin
	return s
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(550)
	f.SetHeight(400)
	f.SetCaption("阴阳历")
	//
	f.label = vcl.NewLabel(f)
	f.label.SetParent(f)
	f.label.SetLeft(10)
	f.label.SetTop(10)
	f.label.SetWidth(40)
	f.label.SetHeight(30)
	f.label.SetTextBuf("输入年月\neg:202101")

	//
	f.year = vcl.NewEdit(f)
	f.year.SetParent(f)
	f.year.SetTop(10)
	f.year.SetLeft(80)
	f.year.SetWidth(60)
	f.year.SetHeight(30)
	f.year.SetTextBuf("202101")

	//
	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetTop(8)
	f.Btn1.SetLeft(140)
	f.Btn1.SetWidth(40)
	f.Btn1.SetHeight(29)
	f.Btn1.SetCaption("OK")
	f.Btn1.SetOnClick(f.OnBtn1Click)
	//
	//f.Btn2 = vcl.NewButton(f)
	//f.Btn2.SetParent(f)
	//f.Btn2.SetBounds(180, 10, 88, 28)
	//f.Btn2.SetCaption("Button2")
	//f.Btn2.SetOnClick(f.OnBtn2Click)
	//---
	t := time.Now().Local()
	y, m, _, h := t.Year(), int(t.Month()), t.Day(), t.Hour()
	//---
	f.edit = vcl.NewEdit(f)
	f.edit.SetParent(f)
	f.edit.SetTop(50)
	f.edit.SetWidth(260)
	f.edit.SetHeight(60)
	f.edit.SetLeft(20)
	f.edit.SetText(info(y, m, 1, h))
	//
	f.edit2 = vcl.NewEdit(f)
	f.edit2.SetParent(f)
	f.edit2.SetTop(70)
	f.edit2.SetWidth(260)
	f.edit2.SetHeight(60)
	f.edit2.SetLeft(20)
	f.edit2.SetText(info(y, m, 2, h))
	//
	f.edit3 = vcl.NewEdit(f)
	f.edit3.SetParent(f)
	f.edit3.SetTop(90)
	f.edit3.SetWidth(260)
	f.edit3.SetHeight(60)
	f.edit3.SetLeft(20)
	f.edit3.SetText(info(y, m, 3, h))
	//
	f.edit4 = vcl.NewEdit(f)
	f.edit4.SetParent(f)
	f.edit4.SetTop(110)
	f.edit4.SetWidth(260)
	f.edit4.SetHeight(60)
	f.edit4.SetLeft(20)
	f.edit4.SetText(info(y, m, 4, h))
	//
	f.edit5 = vcl.NewEdit(f)
	f.edit5.SetParent(f)
	f.edit5.SetTop(130)
	f.edit5.SetWidth(260)
	f.edit5.SetHeight(60)
	f.edit5.SetLeft(20)
	f.edit5.SetText(info(y, m, 5, h))
	//
	f.edit6 = vcl.NewEdit(f)
	f.edit6.SetParent(f)
	f.edit6.SetTop(150)
	f.edit6.SetWidth(260)
	f.edit6.SetHeight(60)
	f.edit6.SetLeft(20)
	f.edit6.SetText(info(y, m, 6, h))
	//
	f.edit7 = vcl.NewEdit(f)
	f.edit7.SetParent(f)
	f.edit7.SetTop(170)
	f.edit7.SetWidth(260)
	f.edit7.SetHeight(60)
	f.edit7.SetLeft(20)
	f.edit7.SetText(info(y, m, 7, h))
	//
	f.edit8 = vcl.NewEdit(f)
	f.edit8.SetParent(f)
	f.edit8.SetTop(190)
	f.edit8.SetWidth(260)
	f.edit8.SetHeight(60)
	f.edit8.SetLeft(20)
	f.edit8.SetText(info(y, m, 8, h))
	//
	f.edit9 = vcl.NewEdit(f)
	f.edit9.SetParent(f)
	f.edit9.SetTop(210)
	f.edit9.SetWidth(260)
	f.edit9.SetHeight(60)
	f.edit9.SetLeft(20)
	f.edit9.SetText(info(y, m, 9, h))
	//
	f.edit10 = vcl.NewEdit(f)
	f.edit10.SetParent(f)
	f.edit10.SetTop(230)
	f.edit10.SetWidth(260)
	f.edit10.SetHeight(60)
	f.edit10.SetLeft(20)
	f.edit10.SetText(info(y, m, 10, h))
	//
	f.edit11 = vcl.NewEdit(f)
	f.edit11.SetParent(f)
	f.edit11.SetTop(250)
	f.edit11.SetWidth(260)
	f.edit11.SetHeight(60)
	f.edit11.SetLeft(20)
	f.edit11.SetText(info(y, m, 11, h))
	//
	f.edit12 = vcl.NewEdit(f)
	f.edit12.SetParent(f)
	f.edit12.SetTop(270)
	f.edit12.SetWidth(260)
	f.edit12.SetHeight(60)
	f.edit12.SetLeft(20)
	f.edit12.SetText(info(y, m, 12, h))
	//
	f.edit13 = vcl.NewEdit(f)
	f.edit13.SetParent(f)
	f.edit13.SetTop(290)
	f.edit13.SetWidth(260)
	f.edit13.SetHeight(60)
	f.edit13.SetLeft(20)
	f.edit13.SetText(info(y, m, 13, h))
	//
	f.edit14 = vcl.NewEdit(f)
	f.edit14.SetParent(f)
	f.edit14.SetTop(310)
	f.edit14.SetWidth(260)
	f.edit14.SetHeight(60)
	f.edit14.SetLeft(20)
	f.edit14.SetText(info(y, m, 14, h))
	//
	f.edit15 = vcl.NewEdit(f)
	f.edit15.SetParent(f)
	f.edit15.SetTop(330)
	f.edit15.SetWidth(260)
	f.edit15.SetHeight(60)
	f.edit15.SetLeft(20)
	f.edit15.SetText(info(y, m, 15, h))
	//-------
	f.edit16 = vcl.NewEdit(f)
	f.edit16.SetParent(f)
	f.edit16.SetTop(50)
	f.edit16.SetWidth(260)
	f.edit16.SetHeight(60)
	f.edit16.SetLeft(270)
	f.edit16.SetText(info(y, m, 16, h))
	//
	f.edit17 = vcl.NewEdit(f)
	f.edit17.SetParent(f)
	f.edit17.SetTop(70)
	f.edit17.SetWidth(260)
	f.edit17.SetHeight(60)
	f.edit17.SetLeft(270)
	f.edit17.SetText(info(y, m, 17, h))
	//
	f.edit18 = vcl.NewEdit(f)
	f.edit18.SetParent(f)
	f.edit18.SetTop(90)
	f.edit18.SetWidth(260)
	f.edit18.SetHeight(60)
	f.edit18.SetLeft(270)
	f.edit18.SetText(info(y, m, 18, h))
	//
	f.edit19 = vcl.NewEdit(f)
	f.edit19.SetParent(f)
	f.edit19.SetTop(110)
	f.edit19.SetWidth(260)
	f.edit19.SetHeight(60)
	f.edit19.SetLeft(270)
	f.edit19.SetText(info(y, m, 19, h))
	//
	f.edit20 = vcl.NewEdit(f)
	f.edit20.SetParent(f)
	f.edit20.SetTop(130)
	f.edit20.SetWidth(260)
	f.edit20.SetHeight(60)
	f.edit20.SetLeft(270)
	f.edit20.SetText(info(y, m, 20, h))
	//
	f.edit21 = vcl.NewEdit(f)
	f.edit21.SetParent(f)
	f.edit21.SetTop(150)
	f.edit21.SetWidth(260)
	f.edit21.SetHeight(60)
	f.edit21.SetLeft(270)
	f.edit21.SetText(info(y, m, 21, h))
	//
	f.edit22 = vcl.NewEdit(f)
	f.edit22.SetParent(f)
	f.edit22.SetTop(170)
	f.edit22.SetWidth(260)
	f.edit22.SetHeight(60)
	f.edit22.SetLeft(270)
	f.edit22.SetText(info(y, m, 22, h))
	//
	f.edit23 = vcl.NewEdit(f)
	f.edit23.SetParent(f)
	f.edit23.SetTop(190)
	f.edit23.SetWidth(260)
	f.edit23.SetHeight(60)
	f.edit23.SetLeft(270)
	f.edit23.SetText(info(y, m, 23, h))
	//
	f.edit24 = vcl.NewEdit(f)
	f.edit24.SetParent(f)
	f.edit24.SetTop(210)
	f.edit24.SetWidth(260)
	f.edit24.SetHeight(60)
	f.edit24.SetLeft(270)
	f.edit24.SetText(info(y, m, 24, h))
	//
	f.edit25 = vcl.NewEdit(f)
	f.edit25.SetParent(f)
	f.edit25.SetTop(230)
	f.edit25.SetWidth(260)
	f.edit25.SetHeight(60)
	f.edit25.SetLeft(270)
	f.edit25.SetText(info(y, m, 25, h))
	//
	f.edit26 = vcl.NewEdit(f)
	f.edit26.SetParent(f)
	f.edit26.SetTop(250)
	f.edit26.SetWidth(260)
	f.edit26.SetHeight(60)
	f.edit26.SetLeft(270)
	f.edit26.SetText(info(y, m, 26, h))
	//
	f.edit27 = vcl.NewEdit(f)
	f.edit27.SetParent(f)
	f.edit27.SetTop(270)
	f.edit27.SetWidth(260)
	f.edit27.SetHeight(60)
	f.edit27.SetLeft(270)
	f.edit27.SetText(info(y, m, 27, h))
	//
	f.edit28 = vcl.NewEdit(f)
	f.edit28.SetParent(f)
	f.edit28.SetTop(290)
	f.edit28.SetWidth(260)
	f.edit28.SetHeight(60)
	f.edit28.SetLeft(270)
	f.edit28.SetText(info(y, m, 28, h))
	//
	f.edit29 = vcl.NewEdit(f)
	f.edit29.SetParent(f)
	f.edit29.SetTop(310)
	f.edit29.SetWidth(260)
	f.edit29.SetHeight(60)
	f.edit29.SetLeft(270)
	f.edit29.SetText(info(y, m, 29, h))
	//
	f.edit30 = vcl.NewEdit(f)
	f.edit30.SetParent(f)
	f.edit30.SetTop(330)
	f.edit30.SetWidth(260)
	f.edit30.SetHeight(60)
	f.edit30.SetLeft(270)
	f.edit30.SetText(info(y, m, 30, h))
	//
	f.edit31 = vcl.NewEdit(f)
	f.edit31.SetParent(f)
	f.edit31.SetTop(350)
	f.edit31.SetWidth(260)
	f.edit31.SetHeight(60)
	f.edit31.SetLeft(270)
	f.edit31.SetText(info(y, m, 31, h))
	//-------------
	if m == 4 || m == 6 || m == 9 || m == 11 {
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
	b := (y%4 == 0 && y%100 != 0) || y%400 == 0
	//闰年2月
	if m == 2 && b == true {
		f.edit30 = vcl.NewEdit(f)
		f.edit30.SetParent(f)
		f.edit30.SetTop(330)
		f.edit30.SetWidth(260)
		f.edit30.SetHeight(60)
		f.edit30.SetLeft(270)
		f.edit30.SetText("")
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
	if m == 2 && b == false {
		f.edit29 = vcl.NewEdit(f)
		f.edit29.SetParent(f)
		f.edit29.SetTop(310)
		f.edit29.SetWidth(260)
		f.edit29.SetHeight(60)
		f.edit29.SetLeft(270)
		f.edit29.SetText("")
		f.edit30 = vcl.NewEdit(f)
		f.edit30.SetParent(f)
		f.edit30.SetTop(330)
		f.edit30.SetWidth(260)
		f.edit30.SetHeight(60)
		f.edit30.SetLeft(270)
		f.edit30.SetText("")
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
}
func (f *TMainForm) show() {
	text := f.year.Text()
	if text == "" {
		vcl.ShowMessage("年月数字为空eg: 202001")
		os.Exit(0)
	}
	years := text[:4]
	months := text[4:]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	m, err := strconv.Atoi(months)
	if err != nil {
		fmt.Println("months", err)
	}

	h := 12
	//---
	f.edit = vcl.NewEdit(f)
	f.edit.SetParent(f)
	f.edit.SetTop(50)
	f.edit.SetWidth(260)
	f.edit.SetHeight(60)
	f.edit.SetLeft(20)
	f.edit.SetText(info(y, m, 1, h))
	//
	f.edit2 = vcl.NewEdit(f)
	f.edit2.SetParent(f)
	f.edit2.SetTop(70)
	f.edit2.SetWidth(260)
	f.edit2.SetHeight(60)
	f.edit2.SetLeft(20)
	f.edit2.SetText(info(y, m, 2, h))
	//
	f.edit3 = vcl.NewEdit(f)
	f.edit3.SetParent(f)
	f.edit3.SetTop(90)
	f.edit3.SetWidth(260)
	f.edit3.SetHeight(60)
	f.edit3.SetLeft(20)
	f.edit3.SetText(info(y, m, 3, h))
	//
	f.edit4 = vcl.NewEdit(f)
	f.edit4.SetParent(f)
	f.edit4.SetTop(110)
	f.edit4.SetWidth(260)
	f.edit4.SetHeight(60)
	f.edit4.SetLeft(20)
	f.edit4.SetText(info(y, m, 4, h))
	//
	f.edit5 = vcl.NewEdit(f)
	f.edit5.SetParent(f)
	f.edit5.SetTop(130)
	f.edit5.SetWidth(260)
	f.edit5.SetHeight(60)
	f.edit5.SetLeft(20)
	f.edit5.SetText(info(y, m, 5, h))
	//
	f.edit6 = vcl.NewEdit(f)
	f.edit6.SetParent(f)
	f.edit6.SetTop(150)
	f.edit6.SetWidth(260)
	f.edit6.SetHeight(60)
	f.edit6.SetLeft(20)
	f.edit6.SetText(info(y, m, 6, h))
	//
	f.edit7 = vcl.NewEdit(f)
	f.edit7.SetParent(f)
	f.edit7.SetTop(170)
	f.edit7.SetWidth(260)
	f.edit7.SetHeight(60)
	f.edit7.SetLeft(20)
	f.edit7.SetText(info(y, m, 7, h))
	//
	f.edit8 = vcl.NewEdit(f)
	f.edit8.SetParent(f)
	f.edit8.SetTop(190)
	f.edit8.SetWidth(260)
	f.edit8.SetHeight(60)
	f.edit8.SetLeft(20)
	f.edit8.SetText(info(y, m, 8, h))
	//
	f.edit9 = vcl.NewEdit(f)
	f.edit9.SetParent(f)
	f.edit9.SetTop(210)
	f.edit9.SetWidth(260)
	f.edit9.SetHeight(60)
	f.edit9.SetLeft(20)
	f.edit9.SetText(info(y, m, 9, h))
	//
	f.edit10 = vcl.NewEdit(f)
	f.edit10.SetParent(f)
	f.edit10.SetTop(230)
	f.edit10.SetWidth(260)
	f.edit10.SetHeight(60)
	f.edit10.SetLeft(20)
	f.edit10.SetText(info(y, m, 10, h))
	//
	f.edit11 = vcl.NewEdit(f)
	f.edit11.SetParent(f)
	f.edit11.SetTop(250)
	f.edit11.SetWidth(260)
	f.edit11.SetHeight(60)
	f.edit11.SetLeft(20)
	f.edit11.SetText(info(y, m, 11, h))
	//
	f.edit12 = vcl.NewEdit(f)
	f.edit12.SetParent(f)
	f.edit12.SetTop(270)
	f.edit12.SetWidth(260)
	f.edit12.SetHeight(60)
	f.edit12.SetLeft(20)
	f.edit12.SetText(info(y, m, 12, h))
	//
	f.edit13 = vcl.NewEdit(f)
	f.edit13.SetParent(f)
	f.edit13.SetTop(290)
	f.edit13.SetWidth(260)
	f.edit13.SetHeight(60)
	f.edit13.SetLeft(20)
	f.edit13.SetText(info(y, m, 13, h))
	//
	f.edit14 = vcl.NewEdit(f)
	f.edit14.SetParent(f)
	f.edit14.SetTop(310)
	f.edit14.SetWidth(260)
	f.edit14.SetHeight(60)
	f.edit14.SetLeft(20)
	f.edit14.SetText(info(y, m, 14, h))
	//
	f.edit15 = vcl.NewEdit(f)
	f.edit15.SetParent(f)
	f.edit15.SetTop(330)
	f.edit15.SetWidth(260)
	f.edit15.SetHeight(60)
	f.edit15.SetLeft(20)
	f.edit15.SetText(info(y, m, 15, h))
	//-------
	f.edit16 = vcl.NewEdit(f)
	f.edit16.SetParent(f)
	f.edit16.SetTop(50)
	f.edit16.SetWidth(260)
	f.edit16.SetHeight(60)
	f.edit16.SetLeft(270)
	f.edit16.SetText(info(y, m, 16, h))
	//
	f.edit17 = vcl.NewEdit(f)
	f.edit17.SetParent(f)
	f.edit17.SetTop(70)
	f.edit17.SetWidth(260)
	f.edit17.SetHeight(60)
	f.edit17.SetLeft(270)
	f.edit17.SetText(info(y, m, 17, h))
	//
	f.edit18 = vcl.NewEdit(f)
	f.edit18.SetParent(f)
	f.edit18.SetTop(90)
	f.edit18.SetWidth(260)
	f.edit18.SetHeight(60)
	f.edit18.SetLeft(270)
	f.edit18.SetText(info(y, m, 18, h))
	//
	f.edit19 = vcl.NewEdit(f)
	f.edit19.SetParent(f)
	f.edit19.SetTop(110)
	f.edit19.SetWidth(260)
	f.edit19.SetHeight(60)
	f.edit19.SetLeft(270)
	f.edit19.SetText(info(y, m, 19, h))
	//
	f.edit20 = vcl.NewEdit(f)
	f.edit20.SetParent(f)
	f.edit20.SetTop(130)
	f.edit20.SetWidth(260)
	f.edit20.SetHeight(60)
	f.edit20.SetLeft(270)
	f.edit20.SetText(info(y, m, 20, h))
	//
	f.edit21 = vcl.NewEdit(f)
	f.edit21.SetParent(f)
	f.edit21.SetTop(150)
	f.edit21.SetWidth(260)
	f.edit21.SetHeight(60)
	f.edit21.SetLeft(270)
	f.edit21.SetText(info(y, m, 21, h))
	//
	f.edit22 = vcl.NewEdit(f)
	f.edit22.SetParent(f)
	f.edit22.SetTop(170)
	f.edit22.SetWidth(260)
	f.edit22.SetHeight(60)
	f.edit22.SetLeft(270)
	f.edit22.SetText(info(y, m, 22, h))
	//
	f.edit23 = vcl.NewEdit(f)
	f.edit23.SetParent(f)
	f.edit23.SetTop(190)
	f.edit23.SetWidth(260)
	f.edit23.SetHeight(60)
	f.edit23.SetLeft(270)
	f.edit23.SetText(info(y, m, 23, h))
	//
	f.edit24 = vcl.NewEdit(f)
	f.edit24.SetParent(f)
	f.edit24.SetTop(210)
	f.edit24.SetWidth(260)
	f.edit24.SetHeight(60)
	f.edit24.SetLeft(270)
	f.edit24.SetText(info(y, m, 24, h))
	//
	f.edit25 = vcl.NewEdit(f)
	f.edit25.SetParent(f)
	f.edit25.SetTop(230)
	f.edit25.SetWidth(260)
	f.edit25.SetHeight(60)
	f.edit25.SetLeft(270)
	f.edit25.SetText(info(y, m, 25, h))
	//
	f.edit26 = vcl.NewEdit(f)
	f.edit26.SetParent(f)
	f.edit26.SetTop(250)
	f.edit26.SetWidth(260)
	f.edit26.SetHeight(60)
	f.edit26.SetLeft(270)
	f.edit26.SetText(info(y, m, 26, h))
	//
	f.edit27 = vcl.NewEdit(f)
	f.edit27.SetParent(f)
	f.edit27.SetTop(270)
	f.edit27.SetWidth(260)
	f.edit27.SetHeight(60)
	f.edit27.SetLeft(270)
	f.edit27.SetText(info(y, m, 27, h))
	//
	f.edit28 = vcl.NewEdit(f)
	f.edit28.SetParent(f)
	f.edit28.SetTop(290)
	f.edit28.SetWidth(260)
	f.edit28.SetHeight(60)
	f.edit28.SetLeft(270)
	f.edit28.SetText(info(y, m, 28, h))
	//
	f.edit29 = vcl.NewEdit(f)
	f.edit29.SetParent(f)
	f.edit29.SetTop(310)
	f.edit29.SetWidth(260)
	f.edit29.SetHeight(60)
	f.edit29.SetLeft(270)
	f.edit29.SetText(info(y, m, 29, h))
	//
	f.edit30 = vcl.NewEdit(f)
	f.edit30.SetParent(f)
	f.edit30.SetTop(330)
	f.edit30.SetWidth(260)
	f.edit30.SetHeight(60)
	f.edit30.SetLeft(270)
	f.edit30.SetText(info(y, m, 30, h))
	//
	f.edit31 = vcl.NewEdit(f)
	f.edit31.SetParent(f)
	f.edit31.SetTop(350)
	f.edit31.SetWidth(260)
	f.edit31.SetHeight(60)
	f.edit31.SetLeft(270)
	f.edit31.SetText(info(y, m, 31, h))
	//-------------
	if m == 4 || m == 6 || m == 9 || m == 11 {
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
	b := (y%4 == 0 && y%100 != 0) || y%400 == 0
	//闰年2月
	if m == 2 && b == true {
		f.edit30 = vcl.NewEdit(f)
		f.edit30.SetParent(f)
		f.edit30.SetTop(330)
		f.edit30.SetWidth(260)
		f.edit30.SetHeight(60)
		f.edit30.SetLeft(270)
		f.edit30.SetText("")
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
	if m == 2 && b == false {
		f.edit29 = vcl.NewEdit(f)
		f.edit29.SetParent(f)
		f.edit29.SetTop(310)
		f.edit29.SetWidth(260)
		f.edit29.SetHeight(60)
		f.edit29.SetLeft(270)
		f.edit29.SetText("")
		f.edit30 = vcl.NewEdit(f)
		f.edit30.SetParent(f)
		f.edit30.SetTop(330)
		f.edit30.SetWidth(260)
		f.edit30.SetHeight(60)
		f.edit30.SetLeft(270)
		f.edit30.SetText("")
		f.edit31 = vcl.NewEdit(f)
		f.edit31.SetParent(f)
		f.edit31.SetTop(350)
		f.edit31.SetWidth(260)
		f.edit31.SetHeight(60)
		f.edit31.SetLeft(270)
		f.edit31.SetText("")
	}
}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
	f.show()

	aboutForm.Show()
}
func (f *TMainForm) OnBtn2Click(sender vcl.IObject) {
	aboutForm.Show()
}

// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("About")
	f.Btn1 = vcl.NewButton(f)
	//f.Btn1.SetName("Btn1")
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("Button1")
	f.Btn1.SetOnClick(f.OnBtn1Click)
	//
	//s := "编辑 编辑 编辑编辑"
	//f.edit = vcl.NewEdit(f)
	//f.edit.SetParent(f)
	//f.edit.SetWidth(120)
	//f.edit.SetHeight(120)
	//f.edit.SetTop(50)
	//f.edit.SetLeft(50)
	//f.edit.SetText(s)
	//
	f.label = vcl.NewLabel(f)
	f.label.SetParent(f)
	f.label.SetLeft(10)
	f.label.SetTop(100)
	f.label.SetWidth(40)
	f.label.SetHeight(60)
	f.label.SetName("dateabout")

}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
	//f.edit.SetText("顶顶顶")
	f.label.SetName("label")
	f.label.SetTextBuf("texttext")
	vcl.ShowMessage("Hello!-------------\n------------------\n")
	//f.Btn1.SetTextBuf("======按钮")
}
