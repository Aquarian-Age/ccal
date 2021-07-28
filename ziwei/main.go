package main

import (
	"fmt"
	"github.com/Aquarian-Age/zwds/ziwei"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strconv"
	"strings"
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(900)
	f.SetHeight(950)
	//f.SetWindowState(types.WsFullScreen)//默认满屏
	f.SetCaption("紫微斗数")
	//f.EnabledMaximize(false)
	//
	f.draw()
	//
	f.OnFormCreateInput(sender)
	//
	f.today()
	//
	f.OnFormAbout()
}

func (f *TMainForm) draw() {
	f.SetOnPaint(func(sender vcl.IObject) {
		canvas := f.Canvas()
		//巳宫
		r := types.TRect{l1, top, r1, b1}
		canvas.FillRect(r)
		//未宫
		r = types.TRect{l2, top, r2, b1}
		canvas.FillRect(r)
		//卯宫
		r = types.TRect{Left: l1, Top: t3, Right: r1, Bottom: b3}
		canvas.FillRect(r)
		//丑宫
		r = types.TRect{l4, t4, r4, b4}
		canvas.FillRect(r)
		//亥宫
		r = types.TRect{l5, t4, r5, b4}
		canvas.FillRect(r)
		//酉宫
		r = types.TRect{l5, t6, r5, b6}
		canvas.FillRect(r)
		//中央
		//r = types.TRect{r1, b1, l5, t4}
		//
		//canvas.FillRect(r)
	})

}

//输入信息
func (f *TMainForm) OnFormCreateInput(vcl.IObject) {
	f.input = vcl.NewEdit(f)
	f.input.SetParent(f)
	f.input.SetTop(b1)
	f.input.SetLeft(r1)
	f.input.SetWidth(120)
	f.input.SetHeight(30)
	f.input.SetTextBuf("2021050413m")

	f.BtnOK = vcl.NewButton(f)
	f.BtnOK.SetParent(f)
	f.BtnOK.SetTop(b1)
	f.BtnOK.SetLeft(r1 + 140)
	f.BtnOK.SetWidth(width)
	f.BtnOK.SetHeight(33)
	f.BtnOK.SetCaption("OK")
	f.BtnOK.SetOnClick(f.OnBtnOK)

}

//btn　show
func (f *TMainForm) OnBtnOK(sender vcl.IObject) {
	text := f.input.Text()
	var y, d, h int
	years := text[:4]
	months := text[4:6]
	days := text[6:8]
	hours := text[8:10]
	sex := text[10:]
	sex = strings.ToLower(sex)

	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	m, err := strconv.Atoi(months)
	if err != nil {
		fmt.Println("months", err)
	}
	d, err = strconv.Atoi(days)
	if err != nil {
		fmt.Println("days", err)
	}
	h, err = strconv.Atoi(hours)
	if err != nil {
		fmt.Println("hours", err)
	}

	if strings.EqualFold(sex, "m") {
		sex = "男"
	}
	if strings.EqualFold(sex, "w") {
		sex = "女"
	}
	zw := ziwei.NewZiWei(y, m, d, h, sex)
	solar := zw.Solar + "性别:" + sex
	gzs := zw.Gzs
	moons := zw.Moon
	mingGong := "命宫:" + zw.MingGong
	shenGong := "身宫:" + zw.ShenGong
	mingZhu := "命主: " + zw.MingZhu
	shenZhu := "身主: " + zw.ShenZhu
	wuxings := zw.WuXing
	yangGui := "阳贵:" + zw.GuiRenYang
	yinGui := "阴贵:" + zw.GuiRenYin
	label1s := solar + "\n" + gzs + "\n" + moons + "  " +
		wuxings + "\n" + mingGong + "  " + mingZhu + "\n" + shenGong + "  " + shenZhu + "\n" + yangGui + " " + yinGui

	riQin := "日禽:" + zw.RiQin
	shiQin := "时禽:" //+ zw.ShiQin
	jianChu := "日建除:" + zw.JianChu
	yueXiang := zw.YueXiang
	jieQi := zw.JieQi
	yueJiang := zw.YueJiang
	label2s := jieQi + "\n" + yueJiang + "\n" + jianChu + "\n" + yueXiang + "\n" + riQin + " " + shiQin
	//
	info := zw.NewInfoZw()
	sihuaY := info.SiHuaY
	sihuaM := info.SiHuaM
	sihuaD := info.SiHuaD
	label3s := sihuaY + "\n" + sihuaM + "\n" + sihuaD
	mgz := zw.YinShou //月干支
	changShengmap := zw.ChangSheng
	var chArr []string
	for i := 0; i < len(mgz); i++ {
		for k, v := range changShengmap {
			if strings.Contains(mgz[i], k) {
				chArr = append(chArr, v)
			}
		}
	}
	mingPanmap := zw.MingPan     //十二宫
	siHuaMap := zw.SiHua         //四化
	star14Map := zw.Star14       //14主星
	miaoMap := zw.Miao           //庙
	yearGanMap := zw.YearGanStar //年干系诸星
	yearZhiMap := zw.YearZhiStar //年支系诸星
	monthStarMap := zw.MonthStar //月系诸星
	hourStarMap := zw.HourStar   //时系诸星
	daXianMap := zw.DaXian       //大限

	f.OnFormInfoShow(label1s, label2s, label3s)
	f.OnFormMgzShow(mgz, chArr)              //月干支　十二长生　btn
	f.OnFormSiHuaShow(siHuaMap)              //四化
	f.OnFormMingGongShow(mingPanmap)         //十二命宫
	f.OnFormStar14Show(star14Map, miaoMap)   //十四主星+庙
	f.OnFormYearGanShow(yearGanMap, miaoMap) //年干系诸星+庙
	f.OnFormYearZhiShow(yearZhiMap)          //年支系诸星
	f.OnFormMonthShow(monthStarMap)          //月系诸星
	f.OnFormHourShow(hourStarMap)            //时系诸星
	f.OnFormDaXianShow(daXianMap)            //大限

}

//abour
func (f *TMainForm) OnFormAbout() {
	f.about = vcl.NewLabel(f)
	f.about.SetParent(f)
	f.about.SetTop(890)
	f.about.SetLeft(l1)
	f.about.Font().SetSize(14)
	f.about.SetTextBuf("开源地址: https://github.com/Aquarian-Age/zwds")
}
