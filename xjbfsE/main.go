package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Aquarian-Age/xa/pkg/gz"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/moon"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TMainForm struct {
	*vcl.TForm
	year                   *vcl.TEdit     //获取输入年数
	month, day, hour       *vcl.TSpinEdit //月日时输入
	BtnOk                  *vcl.TButton
	pgc                    *vcl.TPageControl
	tab, tab1, tab3, about *vcl.TTabSheet
	panelL, panelR         *vcl.TPanel
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

//
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.View(sender)
	f.OnFormShow(sender)
	f.OnFormDefaule(sender)
}

//布局
func (f *TMainForm) View(sender vcl.IObject) {
	f.SetWidth(900)
	f.SetHeight(800)
	f.SetCaption("择日")

	f.pgc = vcl.NewPageControl(f)
	f.pgc.SetParent(f)
	f.pgc.SetAlign(types.AlClient)
	//-------------------
	f.tab = vcl.NewTabSheet(f)
	f.tab.SetPageControl(f.pgc)
	f.tab.SetCaption("择日")

	f.tab1 = vcl.NewTabSheet(f)
	f.tab1.SetPageControl(f.pgc)
	f.tab1.SetCaption("24节气")

	f.tab3 = vcl.NewTabSheet(f)
	f.tab3.SetPageControl(f.pgc)
	f.tab3.SetCaption("月历")

	f.about = vcl.NewTabSheet(f)
	f.about.SetPageControl(f.pgc)
	f.about.SetCaption("About")

	btn := vcl.NewButton(f)
	btn.SetParent(f.about)
	btn.SetCaption("关于")
	btn.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessageFmt(abouts)
	})

}

const abouts = `
协纪辩方书E
v0.6.9
下载地址 https://github.com/Aquarian-Age/ccal/releases
gui开源库: github.com/ying32/govcl
感谢:JetBrains　https://www.jetbrains.com
本程序不含任何担保
`

//获取输入信息
func (f *TMainForm) OnFormShow(sender vcl.IObject) {
	w := int32(300)
	f.year = vcl.NewEdit(f)
	f.year.SetParent(f)
	f.year.SetLeft(w)
	f.year.SetTextBuf("2021")

	w = f.year.Width() + w
	f.month = vcl.NewSpinEdit(f)
	f.month.SetParent(f)
	f.month.SetLeft(w)
	f.month.SetMaxValue(12)
	f.month.SetMinValue(1)

	w = f.month.Width() + w
	f.day = vcl.NewSpinEdit(f)
	f.day.SetParent(f)
	f.day.SetLeft(w)
	f.day.SetMaxValue(31)
	f.day.SetMinValue(1)

	w = f.day.Width() + w
	f.hour = vcl.NewSpinEdit(f)
	f.hour.SetParent(f)
	f.hour.SetLeft(w)
	f.hour.SetMaxValue(23)
	f.hour.SetMinValue(0)

	w = f.hour.Width() + w
	f.BtnOk = vcl.NewButton(f)
	f.BtnOk.SetParent(f)
	f.BtnOk.SetLeft(w)
	f.BtnOk.SetHeight(33)
	f.BtnOk.Font().SetColor(colors.ClRed)
	f.BtnOk.SetCaption("OK")
	f.BtnOk.SetOnClick(f.OnBtnOK)

}

//根据输入信息显示
func (f *TMainForm) OnBtnOK(sender vcl.IObject) {
	f.panelL.Free()
	f.panelL = vcl.NewPanel(f)
	f.panelL.SetParent(f.tab)
	f.panelL.SetWidth(245)
	f.panelL.SetParentBackground(false)
	f.panelL.SetColor(colors.ClWhite)
	f.panelL.Font().SetSize(13)
	f.panelL.Font().SetColor(colors.ClBlack)
	f.panelL.SetAlign(types.AlLeft)

	w := f.panelL.Width()
	f.panelR.Free()
	f.panelR = vcl.NewPanel(f)
	f.panelR.SetParent(f.tab)
	f.panelR.SetLeft(w + 10)
	f.panelR.SetAlign(types.AlClient)

	y, _ := strconv.Atoi(f.year.Text())
	m := int(f.month.Value())
	d := int(f.day.Value())
	h := int(f.hour.Value())
	tx := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	weekN := int(tx.Weekday())
	weeks := WeekName(weekN)
	obj := gz.NewGanZhi(y, m, d, h)
	solars := fmt.Sprintf("%d年%d月%d日%d点 %s\n", y, m, d, h, weeks)
	_, _, _, moons := basic.GetLunar(y, m, d)
	phase := moon.Phase(tx)
	moons = fmt.Sprintf("\n阴历:%s 月相:%5f\n", moons, phase)
	gzs := fmt.Sprintf("%s年%s月%s日%s时\n", obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ)
	nys := fmt.Sprintf("%s\n", obj.GetNaYin())
	ymds := solars + gzs + nys + moons
	jqarr, jqs := obj.Jq24()
	jqs = fmt.Sprintf("%s\n", jqs)
	yjzhi, yjName, yjt := obj.YueJiang()
	yjs := fmt.Sprintf("月将:%s(%s)\n%s\n", yjzhi, yjName, yjt.Format("2006-01-02 15:04:05"))
	huangHei := obj.RiHuangHei()
	huangHei = fmt.Sprintf("\n日黄黑:%s\n", huangHei)
	jianChu := obj.RiJianChu()
	jianChu = fmt.Sprintf("日建除:%s\n", jianChu)
	riQin := obj.RiQin(weekN)
	riQin = fmt.Sprintf("日禽:%s\n", riQin)
	huangHeiH := obj.ShiHuangHei()
	huangHeiH = fmt.Sprintf("时黄黑:%s\n", huangHeiH)
	jqt := obj.Jq24T()
	jueliri := obj.XJBF().JueLiRi(tx, jqt)
	jueliri = fmt.Sprintf("%s\n", jueliri)
	tms, tms1 := objz.GuiDengTianMen() //贵登天门
	if tms != "" || tms1 != "" {
		tms = fmt.Sprintf("贵登天门时:%s %s\n", tms, tms1)
	}
	zhis, xings, chongs, hais := obj.XJBF().TaiSui()
	tsuis := fmt.Sprintf("\n%s\n%s\n%s\n%s\n", zhis, xings, chongs, hais)
	dimus := obj.DimuJing()
	dimus = fmt.Sprintf("\n%s\n", dimus)
	sl := ymds + jqs + yjs + huangHei + jianChu + riQin + huangHeiH + jueliri + tms + tsuis + dimus

	ys, ys1, ys2, ys3 := obj.XJBF().NianBiao()
	s1 := pub.ArrStringToS(ys, " ")
	s2 := pub.ArrStringToS(ys1, " ")
	s3 := pub.ArrStringToS(ys2, " ")
	s4 := pub.ArrStringToS(ys3, " ")

	yzls := obj.XJBF().Yzl()
	arr5, arr6 := obj.XJBF().YueBiao()
	s5 := pub.ArrStringToS(arr5, " ")
	s6 := pub.ArrStringToS(arr6, " ")

	jz60 := gz.GetJzArr()
	arr7, arr8 := obj.XJBF().RiBiao(jz60)
	s7 := pub.ArrStringToS(arr7, " ")
	s8 := pub.ArrStringToS(arr8, " ")

	label := vcl.NewLabel(f)
	label.SetParent(f.panelL)
	label.SetTextBuf(sl)

	suiji := vcl.NewLabel(f)
	suiji.SetParent(f.panelR)
	suiji.SetCaption("岁吉")
	suiji.Font().SetColor(colors.ClRed)
	suiji.Font().SetSize(14)
	left := suiji.Width() + 13
	suijis := vcl.NewLabel(f)
	suijis.SetParent(f.panelR)
	suijis.SetLeft(left)
	suijis.SetCaption(s1)
	top := suiji.Height() + 13
	wugui := vcl.NewLabel(f)
	wugui.SetParent(f.panelR)
	wugui.SetTop(top)
	wugui.SetCaption("五鬼")
	wugui.Font().SetSize(14)
	wugui.Font().SetColor(colors.ClBlack)
	wuguis := vcl.NewLabel(f)
	wuguis.SetParent(f.panelR)
	wuguis.SetLeft(left)
	wuguis.SetTop(top + 3)
	wuguis.SetCaption(s2)
	top += wugui.Height() + 13
	suisha := vcl.NewLabel(f)
	suisha.SetParent(f.panelR)
	suisha.SetTop(top)
	suisha.SetCaption("岁煞")
	suisha.Font().SetSize(14)
	suisha.Font().SetColor(colors.ClYellowgreen)
	suishas := vcl.NewLabel(f)
	suishas.SetParent(f.panelR)
	suishas.SetTop(top + 3)
	suishas.SetLeft(left)
	suishas.SetCaption(s3)
	top += suisha.Height() + 13
	suixiong := vcl.NewLabel(f)
	suixiong.SetParent(f.panelR)
	suixiong.SetTop(top)
	suixiong.SetCaption("岁凶")
	suixiong.Font().SetSize(14)
	suixiong.Font().SetColor(colors.ClBlack)
	suixongs := vcl.NewLabel(f)
	suixongs.SetParent(f.panelR)
	suixongs.SetTop(top + 3)
	suixongs.SetLeft(left)
	suixongs.SetCaption(s4)

	top += suixiong.Height() + 13
	yue := vcl.NewLabel(f)
	yue.SetParent(f.panelR)
	yue.SetCaption("月总论")
	yue.SetTop(top)
	yue.Font().SetSize(14)
	yue.Font().SetColor(colors.ClBlack)
	yues := vcl.NewLabel(f)
	yues.SetParent(f.panelR)
	yues.SetTop(top + 5)
	yues.SetLeft(left + 10)
	yues.SetCaption(yzls)

	top += yues.Top() - 30
	riji := vcl.NewLabel(f)
	riji.SetParent(f.panelR)
	riji.SetCaption("日吉")
	riji.SetTop(top)
	riji.Font().SetSize(14)
	riji.Font().SetColor(colors.ClRed)
	rijis := vcl.NewLabel(f)
	rijis.SetParent(f.panelR)
	rijis.SetTop(top + 5)
	rijis.SetLeft(left)
	rijis.SetCaption(s5)
	top += rijis.Height() + 13
	rixiong := vcl.NewLabel(f)
	rixiong.SetParent(f.panelR)
	rixiong.SetCaption("日凶")
	rixiong.SetTop(top)
	rixiong.Font().SetSize(14)
	rixiong.Font().SetColor(colors.ClBlack)
	rixiongs := vcl.NewLabel(f)
	rixiongs.SetParent(f.panelR)
	rixiongs.SetTop(top + 5)
	rixiongs.SetLeft(left)
	rixiongs.SetCaption(s6)

	top += rixiong.Height() + 23
	shiji := vcl.NewLabel(f)
	shiji.SetParent(f.panelR)
	shiji.SetTop(top)
	shiji.SetCaption("时吉")
	shiji.Font().SetSize(14)
	shiji.Font().SetColor(colors.ClRed)
	shijis := vcl.NewLabel(f)
	shijis.SetParent(f.panelR)
	shijis.SetTop(top + 5)
	shijis.SetLeft(left)
	shijis.SetCaption(s7)
	top += shiji.Height() + 13
	shixiong := vcl.NewLabel(f)
	shixiong.SetParent(f.panelR)
	shixiong.SetTop(top)
	shixiong.SetCaption("时凶")
	shixiong.Font().SetSize(14)
	shixiong.Font().SetColor(colors.ClBlack)
	shixiongs := vcl.NewLabel(f)
	shixiongs.SetParent(f.panelR)
	shixiongs.SetTop(top + 5)
	shixiongs.SetLeft(left)
	shixiongs.SetCaption(s8)

	//tab1 24节气
	var qarr, jarr []string
	for i := 0; i < len(jqarr); i++ {
		if i%2 == 0 {
			qarr = append(qarr, jqarr[i])
		}
		if i%2 == 1 {
			jarr = append(jarr, jqarr[i])
		}
	}
	s9 := pub.ArrStringToS(qarr, "\n")
	jarrs := pub.ArrStringToS(jarr, "\n")
	p24jq := vcl.NewPanel(f)
	p24jq.SetParent(f.tab1)
	p24jq.SetAlign(types.AlClient)

	label24jq := vcl.NewLabel(p24jq)
	label24jq.SetParent(p24jq)
	label24jq.Font().SetSize(13)
	label24jq.SetTextBuf(jarrs)
	label24zq := vcl.NewLabel(p24jq)
	label24zq.SetParent(p24jq)
	label24zq.SetLeft(300)
	label24zq.Font().SetSize(13)
	label24zq.SetTextBuf(s9)

	//tab2 演禽
	// yqobj := qx.NewYanQin(tx.Year(), int(tx.Month()), tx.Day(), tx.Hour())
	// qs0 := fmt.Sprintf("%d年%d月%d日%d时\n%s年%s月%s日%s时\n",
	// 	yqobj.Year, yqobj.Month, yqobj.Day, yqobj.Hour,
	// 	yqobj.Ygz, yqobj.Mgz, yqobj.Dgz, yqobj.Hgz)
	// qs1 := fmt.Sprintf("年禽:%s 月禽:%s 日禽:%s 时禽:%s\n", yqobj.NianQin, yqobj.YueQin, yqobj.RiQin, yqobj.ShiQin)
	// qs2 := fmt.Sprintf("畨禽:%s\n活曜:%s\n", yqobj.FanQin, yqobj.HuoYao)
	// qs3 := fmt.Sprintf("将头:%s\n%s\n%s\n", yqobj.JiangTou, yqobj.FuJiangTa, yqobj.FuJiangWo)
	// qs4 := fmt.Sprintf("%s\n", yqobj.RiQinN)
	// jxs := yqobj.GetSimpleJX()
	// qsjx := fmt.Sprintf("\n%s\n", jxs)
	// yqs := qs0 + qs1 + qs2 + qs3 + qs4 + qsjx

	// pyq := vcl.NewPanel(f)
	// pyq.SetParent(f.tab2)
	// pyq.SetAlign(types.AlClient)
	// lyq := vcl.NewLabel(f)
	// lyq.SetParent(pyq)
	// lyq.SetCaption(yqs)
	// lyq.Font().SetSize(14)
	// lyq.Font().SetColor(colors.ClBlack)

	//tab3 月历
	pyueli := vcl.NewPanel(f)
	pyueli.SetParent(f.tab3)
	pyueli.SetAlign(types.AlClient)
	arr := yueli(tx)
	s10 := pub.ArrStringToS(arr, "\n")
	labelyl := vcl.NewLabel(f)
	labelyl.SetParent(pyueli)
	labelyl.Font().SetSize(13)
	labelyl.SetCaption(s10)

	////鼠标进入显示当前信息
	//f.BtnOk.SetOnMouseEnter(func(sender vcl.IObject) {
	//	f.OnFormDefaule(sender)
	//})
}

//默认显示当前时间
func (f *TMainForm) OnFormDefaule(sender vcl.IObject) {
	f.panelL = vcl.NewPanel(f)
	f.panelL.SetParent(f.tab)
	f.panelL.SetWidth(245)
	f.panelL.SetParentBackground(false)
	f.panelL.SetColor(colors.ClWhite)
	f.panelL.Font().SetSize(13)
	f.panelL.Font().SetColor(colors.ClBlack)
	f.panelL.SetAlign(types.AlLeft)

	w := f.panelL.Width()
	f.panelR = vcl.NewPanel(f)
	f.panelR.SetParent(f.tab)
	f.panelR.SetLeft(w + 10)
	//f.panelR.Font().SetSize(13)
	//f.panelR.Font().SetColor(colors.ClBlack)
	f.panelR.SetAlign(types.AlClient)

	T := time.Now().Local()
	y := T.Year()
	m := int(T.Month())
	d := T.Day()
	h := T.Hour()
	weekN := int(T.Weekday())
	weeks := pub.WeekName(weekN)

	obj := gz.NewGanZhi(y, m, d, h)
	solars := fmt.Sprintf("%d年%d月%d日%d点 %s\n", y, m, d, h, weeks)

	_, _, _, moons := basic.GetLunar(y, m, d)
	phase := moon.Phase(T)
	moons = fmt.Sprintf("\n阴历:%s 月相:%5f\n", moons, phase)
	gzs := fmt.Sprintf("%s年%s月%s日%s时\n", obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ)
	nys := fmt.Sprintf("%s\n", obj.GetNaYin())
	ymds := solars + gzs + nys + moons
	jqarr, jqs := obj.Jq24()
	jqs = fmt.Sprintf("%s\n", jqs)
	yjzhi, yjName, yjt := obj.YueJiang()
	yjs := fmt.Sprintf("月将:%s(%s)\n%s\n", yjzhi, yjName, yjt.Format("2006-01-02 15:04:05"))
	huangHei := obj.RiHuangHei()
	huangHei = fmt.Sprintf("\n日黄黑:%s\n", huangHei)
	jianChu := obj.RiJianChu()
	jianChu = fmt.Sprintf("日建除:%s\n", jianChu)
	riQin := obj.RiQin(weekN)
	riQin = fmt.Sprintf("日禽:%s\n", riQin)
	huangHeiH := obj.ShiHuangHei()
	huangHeiH = fmt.Sprintf("时黄黑:%s\n", huangHeiH)
	jqt := obj.Jq24T()
	jueliri := obj.XJBF().JueLiRi(T, jqt)
	jueliri = fmt.Sprintf("%s\n", jueliri)
	tms, tms1 := obj.GuiDengTianMen()
	if tms != "" || tms1 != "" {
		tms = fmt.Sprintf("贵登天门时:%s %s\n", tms, tms1)
	}
	zhis, xings, chongs, hais := obj.XJBF().TaiSui()
	tsuis := fmt.Sprintf("\n%s\n%s\n%s\n%s\n", zhis, xings, chongs, hais)
	dimus := obj.DimuJing()
	dimus = fmt.Sprintf("\n%s\n", dimus)
	sl := ymds + jqs + yjs + huangHei + jianChu + riQin + huangHeiH + jueliri + tms + tsuis + dimus

	ys, ys1, ys2, ys3 := obj.XJBF().NianBiao()
	s1 := pub.ArrStringToS(ys, " ")
	s2 := pub.ArrStringToS(ys1, " ")
	s3 := pub.ArrStringToS(ys2, " ")
	s4 := pub.ArrStringToS(ys3, " ")

	yzls := obj.XJBF().Yzl()
	arr5, arr6 := obj.XJBF().YueBiao()
	s5 := pub.ArrStringToS(arr5, " ")
	s6 := pub.ArrStringToS(arr6, " ")

	jz60 := gz.GetJzArr()
	arr7, arr8 := obj.XJBF().RiBiao(jz60)
	s7 := pub.ArrStringToS(arr7, " ")
	s8 := pub.ArrStringToS(arr8, " ")

	//tab 纪年信息
	label := vcl.NewLabel(f)
	label.SetParent(f.panelL)
	label.SetTextBuf(sl)

	suiji := vcl.NewLabel(f)
	suiji.SetParent(f.panelR)
	suiji.SetCaption("岁吉")
	suiji.Font().SetColor(colors.ClRed)
	suiji.Font().SetSize(14)
	left := suiji.Width()
	suijis := vcl.NewLabel(f)
	suijis.SetParent(f.panelR)
	suijis.SetLeft(left)
	suijis.SetCaption(s1)
	top := suiji.Height() + 10
	wugui := vcl.NewLabel(f)
	wugui.SetParent(f.panelR)
	wugui.SetTop(top)
	wugui.SetCaption("五鬼")
	wugui.Font().SetSize(14)
	wugui.Font().SetColor(colors.ClBlack)
	wuguis := vcl.NewLabel(f)
	wuguis.SetParent(f.panelR)
	wuguis.SetLeft(left)
	wuguis.SetTop(top)
	wuguis.SetCaption(s2)

	top += wugui.Height() + 13
	suisha := vcl.NewLabel(f)
	suisha.SetParent(f.panelR)
	suisha.SetTop(top)
	suisha.SetCaption("岁煞")
	suisha.Font().SetSize(14)
	suisha.Font().SetColor(colors.ClYellowgreen)
	suishas := vcl.NewLabel(f)
	suishas.SetParent(f.panelR)
	suishas.SetTop(top)
	suishas.SetLeft(left)
	suishas.SetCaption(s3)
	top += suisha.Height() + 13
	suixiong := vcl.NewLabel(f)
	suixiong.SetParent(f.panelR)
	suixiong.SetTop(top)
	suixiong.SetCaption("岁凶")
	suixiong.Font().SetSize(14)
	suixiong.Font().SetColor(colors.ClBlack)
	suixongs := vcl.NewLabel(f)
	suixongs.SetParent(f.panelR)
	suixongs.SetTop(top)
	suixongs.SetLeft(left)
	suixongs.SetCaption(s4)

	top += suixiong.Height() + 13
	yue := vcl.NewLabel(f)
	yue.SetParent(f.panelR)
	yue.SetCaption("月总论")
	yue.SetTop(top)
	yue.Font().SetSize(14)
	yue.Font().SetColor(colors.ClBlack)
	yues := vcl.NewLabel(f)
	yues.SetParent(f.panelR)
	yues.SetTop(top + 13)
	yues.SetLeft(left)
	yues.SetCaption(yzls)

	top += yues.Top() + 13
	riji := vcl.NewLabel(f)
	riji.SetParent(f.panelR)
	riji.SetCaption("日吉")
	riji.SetTop(top)
	riji.Font().SetSize(14)
	riji.Font().SetColor(colors.ClRed)
	rijis := vcl.NewLabel(f)
	rijis.SetParent(f.panelR)
	rijis.SetTop(top + 5)
	rijis.SetLeft(left)
	rijis.SetCaption(s5)
	top += rijis.Height() + 13
	rixiong := vcl.NewLabel(f)
	rixiong.SetParent(f.panelR)
	rixiong.SetCaption("日凶")
	rixiong.SetTop(top)
	rixiong.Font().SetSize(14)
	rixiong.Font().SetColor(colors.ClBlack)
	rixiongs := vcl.NewLabel(f)
	rixiongs.SetParent(f.panelR)
	rixiongs.SetTop(top + 5)
	rixiongs.SetLeft(left)
	rixiongs.SetCaption(s6)

	top += rixiong.Height() + 23
	shiji := vcl.NewLabel(f)
	shiji.SetParent(f.panelR)
	shiji.SetTop(top)
	shiji.SetCaption("时吉")
	shiji.Font().SetSize(14)
	shiji.Font().SetColor(colors.ClRed)
	shijis := vcl.NewLabel(f)
	shijis.SetParent(f.panelR)
	shijis.SetTop(top + 5)
	shijis.SetLeft(left)
	shijis.SetCaption(s7)
	top += shiji.Height() + 13
	shixiong := vcl.NewLabel(f)
	shixiong.SetParent(f.panelR)
	shixiong.SetTop(top)
	shixiong.SetCaption("时凶")
	shixiong.Font().SetSize(14)
	shixiong.Font().SetColor(colors.ClBlack)
	shixiongs := vcl.NewLabel(f)
	shixiongs.SetParent(f.panelR)
	shixiongs.SetTop(top + 5)
	shixiongs.SetLeft(left)
	shixiongs.SetCaption(s8)

	//tab1 24节气
	var qarr, jarr []string
	for i := 0; i < len(jqarr); i++ {
		if i%2 == 0 {
			qarr = append(qarr, jqarr[i])
		}
		if i%2 == 1 {
			jarr = append(jarr, jqarr[i])
		}
	}
	s9 := pub.ArrStringToS(qarr, "\n")
	jarrs := pub.ArrStringToS(jarr, "\n")
	p24jq := vcl.NewPanel(f)
	p24jq.SetParent(f.tab1)
	p24jq.SetAlign(types.AlClient)
	label24jq := vcl.NewLabel(p24jq)
	label24jq.SetParent(p24jq)
	label24jq.Font().SetSize(13)
	label24jq.SetTextBuf(jarrs)
	label24zq := vcl.NewLabel(p24jq)
	label24zq.SetParent(p24jq)
	label24zq.SetLeft(300)
	label24zq.Font().SetSize(13)
	label24zq.SetTextBuf(s9)

	//tab2 演禽
	// yqobj := qx.NewYanQin(T.Year(), int(T.Month()), T.Day(), T.Hour())
	// qs0 := fmt.Sprintf("%d年%d月%d日%d时\n%s年%s月%s日%s时\n",
	// 	yqobj.Year, yqobj.Month, yqobj.Day, yqobj.Hour,
	// 	yqobj.Ygz, yqobj.Mgz, yqobj.Dgz, yqobj.Hgz)
	// qs1 := fmt.Sprintf("年禽:%s 月禽:%s 日禽:%s 时禽:%s\n", yqobj.NianQin, yqobj.YueQin, yqobj.RiQin, yqobj.ShiQin)
	// qs2 := fmt.Sprintf("畨禽:%s\n活曜:%s\n", yqobj.FanQin, yqobj.HuoYao)
	// qs3 := fmt.Sprintf("将头:%s\n%s\n%s\n", yqobj.JiangTou, yqobj.FuJiangTa, yqobj.FuJiangWo)
	// qs4 := fmt.Sprintf("%s\n", yqobj.RiQinN)
	// jxs := yqobj.GetSimpleJX()
	// qsjx := fmt.Sprintf("\n%s\n", jxs)
	// yqs := qs0 + qs1 + qs2 + qs3 + qs4 + qsjx

	// pyq := vcl.NewPanel(f)
	// pyq.SetParent(f.tab2)
	// pyq.SetAlign(types.AlClient)
	// lyq := vcl.NewLabel(f)
	// lyq.SetParent(pyq)
	// lyq.SetCaption(yqs)
	// lyq.Font().SetSize(14)
	// lyq.Font().SetColor(colors.ClBlack)

	//tab3 月历 月历的for循环影响启动速度　这里加个btn缓冲
	pyueli := vcl.NewPanel(f)
	pyueli.SetParent(f.tab3)
	pyueli.SetAlign(types.AlClient)
	arr := yueli(T)
	s10 := pub.ArrStringToS(arr, "\n")
	labelyl := vcl.NewLabel(f)
	labelyl.SetParent(pyueli)
	labelyl.Font().SetSize(13)
	labelyl.SetCaption(s10)

}

func yueli(t time.Time) []string {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	solar := fmt.Sprintf("%d月%d日", month, day)
	_, _, _, moon := basic.GetLunar(year, month, day) //阴历
	mgz := gz.GetMonthGZ(year, month, day, hour)
	dgz := gz.GetDayGZ(year, month, day)
	jianChu := gz.GetRiJianChu(mgz, dgz)   //日建除
	huangHei := gz.GetRiHuangHei(mgz, dgz) //日黄黑
	wd := int(t.Weekday())
	riQin := gz.GetRiQin(wd, dgz)

	b := (year%4 == 0 && year%100 != 0) || year%400 == 0

	var arr []string
	switch month {
	case 4, 6, 9, 11:
		for i := 1; i <= 30; i++ {
			solar = fmt.Sprintf("%d月%d日", month, i)
			_, _, _, moon = basic.GetLunar(year, month, i)
			dgz = gz.GetDayGZ(year, month, i)
			t = time.Date(t.Year(), time.Month(month), i, 0, 0, 0, 0, time.Local)
			mgz = gz.GetMonthGZ(t.Year(), month, i, 22)
			jianChu = gz.GetRiJianChu(mgz, dgz)
			huangHei = gz.GetRiHuangHei(mgz, dgz)
			wd = int(t.Weekday())
			riQin = gz.GetRiQin(wd, dgz)
			s := fmt.Sprintf("%s %s %s %s %s %s", solar, moon, dgz, jianChu, huangHei, riQin)
			arr = append(arr, s)
		}
	case 1, 3, 5, 7, 8, 10, 12:
		for i := 1; i <= 31; i++ {
			solar = fmt.Sprintf("%d月%d日", month, i)
			_, _, _, moon = basic.GetLunar(year, month, i)
			dgz = gz.GetDayGZ(year, month, i)
			t = time.Date(t.Year(), time.Month(month), i, 0, 0, 0, 0, time.Local)
			mgz = gz.GetMonthGZ(t.Year(), month, i, 22)
			jianChu = gz.GetRiJianChu(mgz, dgz)
			huangHei = gz.GetRiHuangHei(mgz, dgz)
			wd = int(t.Weekday())
			riQin = gz.GetRiQin(wd, dgz)
			s := fmt.Sprintf("%s %s %s %s %s %s", solar, moon, dgz, jianChu, huangHei, riQin)
			arr = append(arr, s)
		}
	case 2:
		if month == 2 && b == true {
			for i := 1; i <= 29; i++ {
				solar = fmt.Sprintf("%d月%d日", month, i)
				_, _, _, moon = basic.GetLunar(year, month, i)
				dgz = gz.GetDayGZ(year, month, i)
				t = time.Date(t.Year(), time.Month(month), i, 0, 0, 0, 0, time.Local)
				mgz = gz.GetMonthGZ(t.Year(), month, i, 22)
				jianChu = gz.GetRiJianChu(mgz, dgz)
				huangHei = gz.GetRiHuangHei(mgz, dgz)
				wd = int(t.Weekday())
				riQin = gz.GetRiQin(wd, dgz)
				s := fmt.Sprintf("%s %s %s %s %s %s", solar, moon, dgz, jianChu, huangHei, riQin)
				arr = append(arr, s)
			}
		}
		if month == 2 && b == false {
			for i := 1; i <= 28; i++ {
				solar = fmt.Sprintf("%d月%d日", month, i)
				_, _, _, moon = basic.GetLunar(year, month, i)
				dgz = gz.GetDayGZ(year, month, i)
				t = time.Date(t.Year(), time.Month(month), i, 0, 0, 0, 0, time.Local)
				mgz = gz.GetMonthGZ(t.Year(), month, i, 22)
				jianChu = gz.GetRiJianChu(mgz, dgz)
				huangHei = gz.GetRiHuangHei(mgz, dgz)
				wd = int(t.Weekday())
				riQin = gz.GetRiQin(wd, dgz)
				s := fmt.Sprintf("%s %s %s %s %s %s", solar, moon, dgz, jianChu, huangHei, riQin)
				arr = append(arr, s)
			}
		}
	}
	return arr
}
