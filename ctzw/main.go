package main

import (
	"fmt"
	"github.com/Aquarian-Age/zwds/ctzw"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"strconv"
	"strings"
)

type TMainForm struct {
	*vcl.TForm
	input            *vcl.TEdit //获取输入年数
	month, day, hour *vcl.TSpinEdit
	sex              *vcl.TComboBox
	//
	labely, labelm, labeld, labelh, labelsex *vcl.TLabel //输入框侧边说明
	BtnOK                                    *vcl.TButton
	//
	hail, zil, choul, yinl, maol, chenl, sil, wul, weil, shenl, youl, xul *vcl.TLabel //１２宫
	label2                                                                *vcl.TLabel //显示中间部分
	label3                                                                *vcl.TLabel //宫位顺序说明
}

const (
	top   = 10
	width = 60
)

//画
const (
	//巳宫
	l1 = 10
	r1 = 220
	//未宫
	l2 = 2 * r1
	r2 = 3 * r1
	b1 = 215
	//卯宫
	t3 = top + 2*b1
	b3 = 3 * b1
	//丑宫
	t4 = top + 3*b1
	l4 = r1
	r4 = l4 + b1
	b4 = 4 * b1
	//亥宫
	l5 = r2
	r5 = l5 + b1
	//酉宫
	t6 = b1
	b6 = t3
)

var mainForm *TMainForm

func main() {
	vcl.RunApp(&mainForm)
}
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(900)
	f.SetHeight(900)
	f.SetCaption("策天紫微")
	f.Font().SetSize(14)

	f.draw()
	//
	f.OnFormCreatG12()
	//
	f.OnFormGet(sender)
	//
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
	})

}
func (f *TMainForm) OnFormGet(sender vcl.IObject) {
	f.input = vcl.NewEdit(f)
	f.input.SetParent(f)
	f.input.SetTop(b1)
	f.input.SetLeft(r1)
	f.input.SetWidth(width)
	f.input.SetHeight(30)
	f.input.SetTextBuf("2021")
	//
	f.labely = vcl.NewLabel(f)
	f.labely.SetParent(f)
	f.labely.SetTop(b1)
	f.labely.SetLeft(r1 + width + l1)
	f.labely.SetCaption("阳历年")
	//
	f.month = vcl.NewSpinEdit(f)
	f.month.SetParent(f)
	f.month.SetTop(b1 + 30)
	f.month.SetLeft(r1)
	f.month.SetWidth(width)
	f.month.SetMaxValue(12)
	f.month.SetMinValue(1)

	//
	f.labelm = vcl.NewLabel(f)
	f.labelm.SetParent(f)
	f.labelm.SetTop(b1 + 30)
	f.labelm.SetLeft(r1 + width + l1)
	f.labelm.SetCaption("阳历月")
	//
	f.day = vcl.NewSpinEdit(f)
	f.day.SetParent(f)
	f.day.SetTop(b1 + 30*2)
	f.day.SetLeft(r1)
	f.day.SetWidth(width)
	f.day.SetMaxValue(31)
	f.day.SetMinValue(1)

	//
	f.labeld = vcl.NewLabel(f)
	f.labeld.SetParent(f)
	f.labeld.SetTop(b1 + 30*2)
	f.labeld.SetLeft(r1 + width + l1)
	f.labeld.SetCaption("阳历日")
	//
	f.hour = vcl.NewSpinEdit(f)
	f.hour.SetParent(f)
	f.hour.SetTop(b1 + 30*3)
	f.hour.SetLeft(r1)
	f.hour.SetWidth(width)
	f.hour.SetMaxValue(24)
	f.hour.SetMinValue(0)

	//
	f.labelh = vcl.NewLabel(f)
	f.labelh.SetParent(f)
	f.labelh.SetTop(b1 + 30*3)
	f.labelh.SetLeft(r1 + width + l1)
	f.labelh.SetCaption("阳历时")
	//
	f.sex = vcl.NewComboBox(f)
	f.sex.SetParent(f)
	f.sex.SetTop(b1 + 30*4)
	f.sex.SetLeft(r1)
	f.sex.SetWidth(width + 20)
	f.sex.SetName("sex")
	f.sex.Items().Add("男") //0
	f.sex.Items().Add("女") //1
	f.sex.Items().BeginUpdate()
	f.sex.SetOnSelect(func(sender vcl.IObject) {})
	//
	f.labelsex = vcl.NewLabel(f)
	f.labelsex.SetParent(f)
	f.labelsex.SetTop(b1 + 30*4)
	f.labelsex.SetLeft(r1 + width + 25)
	f.labelsex.SetCaption("性别")

	//
	f.BtnOK = vcl.NewButton(f)
	f.BtnOK.SetParent(f)
	f.BtnOK.SetTop(b1 + 37*5)
	f.BtnOK.SetLeft(r1)
	f.BtnOK.SetWidth(width)
	f.BtnOK.SetHeight(33)
	f.BtnOK.Font().SetColor(colors.ClRed)
	f.BtnOK.SetCaption("OK")
	f.BtnOK.SetOnClick(f.OnBtnOK)
	//-------------------
	text := f.input.Text()
	years := text[:4]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	m := int(f.month.Value())
	d := int(f.day.Value())
	h := int(f.hour.Value())

	index := f.sex.ItemIndex()
	var sex string
	if index == 0 {
		sex = "男"
	}
	if index == 1 {
		sex = "女"
	}
	f.OnFormCreateDefault(y, m, d, h, sex)

}

func (f *TMainForm) OnBtnOK(sender vcl.IObject) {
	text := f.input.Text()
	years := text[:4]
	y, err := strconv.Atoi(years)
	if err != nil {
		fmt.Println("years", err)
	}
	m := int(f.month.Value())
	d := int(f.day.Value())
	h := int(f.hour.Value())

	index := f.sex.ItemIndex()
	var sex string
	if index == 0 {
		sex = "男"
	}
	if index == 1 {
		sex = "女"
	}
	//
	f.OnFormShow(y, m, d, h, sex)

}
func (f *TMainForm) OnFormCreatG12() {
	const height1 = 6 * 30
	const font = 16
	const color = colors.ClBlack
	const step = 190
	f.hail = vcl.NewLabel(f)
	f.hail.SetParent(f)
	f.hail.SetTop(t4 + height1 - l1)
	f.hail.SetLeft(l5 + step)
	f.hail.Font().SetSize(font)
	f.hail.Font().SetColor(color)
	f.hail.SetTextBuf("亥")

	f.zil = vcl.NewLabel(f)
	f.zil.SetParent(f)
	f.zil.SetTop(t4 + height1 - l1)
	f.zil.SetLeft(r4 + step)
	f.zil.Font().SetSize(font)
	f.zil.Font().SetColor(color)
	f.zil.SetTextBuf("子")

	f.choul = vcl.NewLabel(f)
	f.choul.SetParent(f)
	f.choul.SetTop(t4 + height1 - l1)
	f.choul.SetLeft(l4 + step)
	f.choul.Font().SetSize(font)
	f.choul.Font().SetColor(color)
	f.choul.SetTextBuf("丑")

	f.yinl = vcl.NewLabel(f)
	f.yinl.SetParent(f)
	f.yinl.SetTop(t4 + height1 - l1)
	f.yinl.SetLeft(l1 + step)
	f.yinl.Font().SetSize(font)
	f.yinl.Font().SetColor(color)
	f.yinl.SetTextBuf("寅")

	f.maol = vcl.NewLabel(f)
	f.maol.SetParent(f)
	f.maol.SetTop(t3 + height1)
	f.maol.SetLeft(l1 + step)
	f.maol.Font().SetSize(font)
	f.maol.Font().SetColor(color)
	f.maol.SetTextBuf("卯")

	f.chenl = vcl.NewLabel(f)
	f.chenl.SetParent(f)
	f.chenl.SetTop(t6 + height1 + l1)
	f.chenl.SetLeft(l1 + step)
	f.chenl.Font().SetSize(font)
	f.chenl.Font().SetColor(color)
	f.chenl.SetTextBuf("辰")

	f.sil = vcl.NewLabel(f)
	f.sil.SetParent(f)
	f.sil.SetTop(top + height1 - l1)
	f.sil.SetLeft(l1 + step)
	f.sil.Font().SetSize(font)
	f.sil.Font().SetColor(color)
	f.sil.SetTextBuf("巳")

	f.wul = vcl.NewLabel(f)
	f.wul.SetParent(f)
	f.wul.SetTop(top + height1 - l1)
	f.wul.SetLeft(r1 + step)
	f.wul.Font().SetSize(font)
	f.wul.Font().SetColor(color)
	f.wul.SetTextBuf("午")

	f.weil = vcl.NewLabel(f)
	f.weil.SetParent(f)
	f.weil.SetTop(top + height1 - l1)
	f.weil.SetLeft(l2 + step)
	f.weil.Font().SetSize(font)
	f.weil.Font().SetColor(color)
	f.weil.SetTextBuf("未")

	f.shenl = vcl.NewLabel(f)
	f.shenl.SetParent(f)
	f.shenl.SetTop(top + height1 - l1)
	f.shenl.SetLeft(l5 + step)
	f.shenl.Font().SetSize(font)
	f.shenl.Font().SetColor(color)
	f.shenl.SetTextBuf("申")

	f.youl = vcl.NewLabel(f)
	f.youl.SetParent(f)
	f.youl.SetTop(t6 + height1 + l1)
	f.youl.SetLeft(l5 + step)
	f.youl.Font().SetSize(font)
	f.youl.Font().SetColor(color)
	f.youl.SetTextBuf("酉")

	f.xul = vcl.NewLabel(f)
	f.xul.SetParent(f)
	f.xul.SetTop(t3 + height1)
	f.xul.SetLeft(l5 + step)
	f.xul.Font().SetSize(font)
	f.xul.Font().SetColor(color)
	f.xul.SetTextBuf("戌")
}

//default
func (f *TMainForm) OnFormCreateDefault(y, m, d, h int, sex string) {
	ct := ctzw.NewCTZW(y, m, d, h, sex)
	g12 := ct.NewG12()
	g12 = g12.G12Miao()

	hai := g12.Hai
	zi := g12.Zi
	chou := g12.Chou
	yin := g12.Yin
	mao := g12.Mao
	chen := g12.Chen
	si := g12.Si
	wu := g12.Wu
	wei := g12.Wei
	shen := g12.Shen
	you := g12.You
	xu := g12.Xu
	hais := arrToS(hai, "\n")
	zis := arrToS(zi, "\n")
	chous := arrToS(chou, "\n")
	yins := arrToS(yin, "\n")
	maos := arrToS(mao, "\n")
	chens := arrToS(chen, "\n")
	sis := arrToS(si, "\n")
	wus := arrToS(wu, "\n")
	weis := arrToS(wei, "\n")
	shens := arrToS(shen, "\n")
	yous := arrToS(you, "\n")
	xus := arrToS(xu, "\n")
	//
	solar := fmt.Sprintf("%d年%d月%d日%d点 %s", y, m, d, h, sex)
	gzs := fmt.Sprintf("%s年%s月%s日%s时", ct.Ygz, ct.Mgz, ct.Dgz, ct.Hgz)
	moons := ct.Moons
	msg := fmt.Sprintf("命宫:%s 身宫:%s", ct.Mg, ct.Sg)
	label2s := solar + "\n" + gzs + "\n" + moons + "\n\n" + msg
	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f)
	f.label2.SetTop(b1 + 10)
	f.label2.SetLeft(2 * r1)
	f.label2.SetWidth(r1)
	f.label2.Font().SetSize(12)
	f.label2.Font().SetColor(colors.ClBlack)
	f.label2.SetTextBuf(label2s)
	//宫位顺序说明
	label3s := "紫微 天虚 天贵 天印 天寿 天空 红鸾 天库 天贯 文昌 天福 天禄\n" +
		"命宫 财帛 兄弟 田宅 男女 奴仆 妻妾 疾厄 迁移 官禄 福德 相貌\n" +
		"天杖 天异 毛头 天刃\n" +
		"四行 大限\n" +
		"五行 小限\n" +
		"天哭 天刑天姚 三台八座 龙池鳳阁\n\n" +
		"下载 https://github.com/Aquarian-Age/zwds"
	f.label3 = vcl.NewLabel(f)
	f.label3.SetParent(f)
	f.label3.SetTop(2*b1 + 10)
	f.label3.SetLeft(r1)
	f.label3.SetWidth(r1)
	f.label3.SetTextBuf(label3s)
	f.label3.Font().SetSize(11)
	f.label3.Font().SetColor(colors.ClBlack)
	//
	f.hail = vcl.NewLabel(f)
	f.hail.SetParent(f)
	f.hail.SetTop(t4)
	f.hail.SetLeft(l5)
	f.hail.SetTextBuf(hais)

	f.zil = vcl.NewLabel(f)
	f.zil.SetParent(f)
	f.zil.SetTop(t4)
	f.zil.SetLeft(r4)
	f.zil.SetTextBuf(zis)

	f.choul = vcl.NewLabel(f)
	f.choul.SetParent(f)
	f.choul.SetTop(t4)
	f.choul.SetLeft(l4)
	f.choul.SetTextBuf(chous)

	f.yinl = vcl.NewLabel(f)
	f.yinl.SetParent(f)
	f.yinl.SetTop(t4)
	f.yinl.SetLeft(l1)
	f.yinl.SetTextBuf(yins)

	f.maol = vcl.NewLabel(f)
	f.maol.SetParent(f)
	f.maol.SetTop(t3)
	f.maol.SetLeft(l1)
	f.maol.SetTextBuf(maos)

	f.chenl = vcl.NewLabel(f)
	f.chenl.SetParent(f)
	f.chenl.SetTop(t6)
	f.chenl.SetLeft(l1)
	f.chenl.SetTextBuf(chens)

	f.sil = vcl.NewLabel(f)
	f.sil.SetParent(f)
	f.sil.SetTop(top)
	f.sil.SetLeft(l1)
	f.sil.SetTextBuf(sis)

	f.wul = vcl.NewLabel(f)
	f.wul.SetParent(f)
	f.wul.SetTop(top)
	f.wul.SetLeft(r1)
	f.wul.SetTextBuf(wus)

	f.weil = vcl.NewLabel(f)
	f.weil.SetParent(f)
	f.weil.SetTop(top)
	f.weil.SetLeft(l2)
	f.weil.SetTextBuf(weis)

	f.shenl = vcl.NewLabel(f)
	f.shenl.SetParent(f)
	f.shenl.SetTop(top)
	f.shenl.SetLeft(l5)
	f.shenl.SetTextBuf(shens)

	f.youl = vcl.NewLabel(f)
	f.youl.SetParent(f)
	f.youl.SetTop(t6)
	f.youl.SetLeft(l5)
	f.youl.SetTextBuf(yous)

	f.xul = vcl.NewLabel(f)
	f.xul.SetParent(f)
	f.xul.SetTop(t3)
	f.xul.SetLeft(l5)
	f.xul.SetTextBuf(xus)

}

//------------------------------------------------------------

//show
func (f *TMainForm) OnFormShow(y, m, d, h int, sex string) {
	ct := ctzw.NewCTZW(y, m, d, h, sex)
	g12 := ct.NewG12()
	g12 = g12.G12Miao()

	hai := g12.Hai
	zi := g12.Zi
	chou := g12.Chou
	yin := g12.Yin
	mao := g12.Mao
	chen := g12.Chen
	si := g12.Si
	wu := g12.Wu
	wei := g12.Wei
	shen := g12.Shen
	you := g12.You
	xu := g12.Xu
	hais := arrToS(hai, "\n")
	zis := arrToS(zi, "\n")
	chous := arrToS(chou, "\n")
	yins := arrToS(yin, "\n")
	maos := arrToS(mao, "\n")
	chens := arrToS(chen, "\n")
	sis := arrToS(si, "\n")
	wus := arrToS(wu, "\n")
	weis := arrToS(wei, "\n")
	shens := arrToS(shen, "\n")
	yous := arrToS(you, "\n")
	xus := arrToS(xu, "\n")
	//-----------------
	solar := fmt.Sprintf("%d年%d月%d日%d点 %s", y, m, d, h, sex)
	gzs := fmt.Sprintf("%s年%s月%s日%s时", ct.Ygz, ct.Mgz, ct.Dgz, ct.Hgz)
	moons := ct.Moons
	msg := fmt.Sprintf("命宫:%s 身宫:%s", ct.Mg, ct.Sg)
	label2s := solar + "\n" + gzs + "\n" + moons + "\n\n" + msg
	f.label2.Free()
	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f)
	f.label2.SetTop(b1)
	f.label2.SetLeft(2 * r1)
	f.label2.SetWidth(r1)
	f.label2.Font().SetSize(12)
	f.label2.Font().SetColor(colors.ClBlack)
	f.label2.SetTextBuf(label2s)
	//
	f.hail.Free()
	f.hail = vcl.NewLabel(f)
	f.hail.SetParent(f)
	f.hail.SetTop(t4)
	f.hail.SetLeft(l5)
	f.hail.SetTextBuf(hais)
	f.zil.Free()
	f.zil = vcl.NewLabel(f)
	f.zil.SetParent(f)
	f.zil.SetTop(t4)
	f.zil.SetLeft(r4)
	f.zil.SetTextBuf(zis)
	f.choul.Free()
	f.choul = vcl.NewLabel(f)
	f.choul.SetParent(f)
	f.choul.SetTop(t4)
	f.choul.SetLeft(l4)
	f.choul.SetTextBuf(chous)
	f.yinl.Free()
	f.yinl = vcl.NewLabel(f)
	f.yinl.SetParent(f)
	f.yinl.SetTop(t4)
	f.yinl.SetLeft(l1)
	f.yinl.SetTextBuf(yins)
	f.maol.Free()
	f.maol = vcl.NewLabel(f)
	f.maol.SetParent(f)
	f.maol.SetTop(t3)
	f.maol.SetLeft(l1)
	f.maol.SetTextBuf(maos)
	f.chenl.Free()
	f.chenl = vcl.NewLabel(f)
	f.chenl.SetParent(f)
	f.chenl.SetTop(t6)
	f.chenl.SetLeft(l1)
	f.chenl.SetTextBuf(chens)
	f.sil.Free()
	f.sil = vcl.NewLabel(f)
	f.sil.SetParent(f)
	f.sil.SetTop(top)
	f.sil.SetLeft(l1)
	f.sil.SetTextBuf(sis)
	f.wul.Free()
	f.wul = vcl.NewLabel(f)
	f.wul.SetParent(f)
	f.wul.SetTop(top)
	f.wul.SetLeft(r1)
	f.wul.SetTextBuf(wus)
	f.weil.Free()
	f.weil = vcl.NewLabel(f)
	f.weil.SetParent(f)
	f.weil.SetTop(top)
	f.weil.SetLeft(l2)
	f.weil.SetTextBuf(weis)
	f.shenl.Free()
	f.shenl = vcl.NewLabel(f)
	f.shenl.SetParent(f)
	f.shenl.SetTop(top)
	f.shenl.SetLeft(l5)
	f.shenl.SetTextBuf(shens)
	f.youl.Free()
	f.youl = vcl.NewLabel(f)
	f.youl.SetParent(f)
	f.youl.SetTop(t6)
	f.youl.SetLeft(l5)
	f.youl.SetTextBuf(yous)

	f.xul.Free()
	f.xul = vcl.NewLabel(f)
	f.xul.SetParent(f)
	f.xul.SetTop(t3)
	f.xul.SetLeft(l5)
	f.xul.SetTextBuf(xus)

}
//数组转字符串
func arrToS(arr []string, want string) string {
	return strings.Join(arr, want)
}