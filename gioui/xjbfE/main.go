package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gonoto/notosans"
	"github.com/starainrt/astro/basic"
	"github.com/starainrt/astro/moon"
	"liangzi.local/qx/pkg/gz"
	"liangzi.local/qx/pkg/pub"
)

var (
	t       time.Time
	lyName  string
	ly1Name string
	ly2Name string
	ly3Name string
	lmName  string
	lm1Name string
	ldName  string
	ld1Name string

	lytxt  string //岁吉
	ly1txt string //岁煞
	ly2txt string //五鬼
	ly3txt string //岁凶

	lmtxt  string //月吉
	lm1txt string //月凶

	ldtxt  string //日吉
	ld1txt string //日凶

	jqs, zqs string //24节气

	btn3s  string //基本纪年信息
	btnArr []string

	btn4s string //关于
)

type UI struct {
	th *material.Theme //主题字体
	//input materialx.InputLayoutStyle //输入信息
	input            widget.Editor
	btn1, btn2, btn3 widget.Clickable
	btn4             widget.Clickable //关于
}

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("协纪辩方书E"),
			app.Size(unit.Dp(360), unit.Dp(640)),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	ui := newUI()
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			gtx.Constraints.Min = image.Pt(140, 40)
			if ui.btn1.Clicked() {
				cleanAll()
				cleanjq24()
				input := ui.input.Text()
				if len(input) == 10 {
					t = convF(input)
					addName()
					xjbfF(t)
				} else {
					t = gtx.Now.Local()
					addName()
					xjbfF(t)
				}
			}
			if ui.btn2.Clicked() {
				cleanAll()
				input := ui.input.Text()
				if len(input) == 10 {
					t = convF(input)
					jq24F(t)
				} else {
					t = gtx.Now.Local()
					jq24F(t)
				}
			}
			if ui.btn3.Clicked() {
				cleanAll()
				cleanjq24()
				input := ui.input.Text()
				if len(input) == 10 {
					t = convF(input)
					btn3F(t)
				} else {
					t = gtx.Now.Local()
					btn3F(t)
				}
			}
			if ui.btn4.Clicked() {
				cleanAll()
				cleanjq24()
				btn4F()
			}

			ui.btnLayout(gtx)
			ui.lyLayout(gtx)
			ui.ly1Layout(gtx)
			ui.ly2Layout(gtx)
			ui.ly3Layout(gtx)

			ui.lmLayout(gtx)
			ui.lm1Layout(gtx)
			ui.ldLayout(gtx)
			ui.ld1Layout(gtx)

			ui.jq24Layout(gtx)
			ui.btn3Layout(gtx)
			ui.btn4Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
	return nil
}

func newUI() *UI {
	fonts := gofont.Collection()
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(fonts)

	input := widget.Editor{}
	input.SingleLine = true

	btn1 := widget.Clickable{}
	btn2 := widget.Clickable{}
	btn3 := widget.Clickable{}
	btn4 := widget.Clickable{}

	return &UI{
		th,
		input,
		btn1,
		btn2,
		btn3,
		btn4,
	}
}

//中文字体
func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

//input btn 布局
func (ui *UI) btnLayout(gtx layout.Context) layout.Dimensions {
	space1 := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	borderWidth := float32(0.3)
	borderColor := color.NRGBA{A: 107}
	if ui.input.Focused() {
		borderColor = ui.th.Palette.ContrastBg
		borderWidth = 3
	}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(borderWidth),
	}

	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(2)).Layout(gtx,
					material.Editor(ui.th, &ui.input, "2021010203").Layout)
			})
		}),
		space1,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn1 := material.Button(ui.th, &ui.btn1, "择日")
			btn1.Font = text.Font{Typeface: "Noto"}
			btn1.Background = color.NRGBA{30, 144, 255, 200}
			btn1.Font.Style = 12
			return btn1.Layout(gtx)
		}),
		space1,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn2 := material.Button(ui.th, &ui.btn2, "节气")
			btn2.Font = text.Font{Typeface: "Noto"}
			btn2.Background = color.NRGBA{30, 144, 255, 200}
			return btn2.Layout(gtx)
		}),
		space1,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn3l := material.Button(ui.th, &ui.btn3, "纪年")
			btn3l.Font = text.Font{Typeface: "Noto"}
			btn3l.Background = color.NRGBA{30, 144, 255, 200}
			return btn3l.Layout(gtx)
		}),
		space1,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn4 := material.Button(ui.th, &ui.btn4, "关于")
			btn4.Font = text.Font{Typeface: "Noto"}
			btn4.Background = color.NRGBA{30, 144, 255, 200}
			return btn4.Layout(gtx)
		}),
	)
}

//岁吉
func (ui *UI) lyLayout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(45),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(40)}.Layout)

	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				ly := material.H6(ui.th, lyName)
				ly.Font = text.Font{Typeface: "Noto"}
				ly.Color = color.NRGBA{255, 0, 0, 255} //红色
				return ly.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lybody := material.Body1(ui.th, lytxt)
				lybody.Font = text.Font{Typeface: "Noto"}
				lybody.Color = color.NRGBA{0, 0, 0, 255}
				return lybody.Layout(gtx)
			}),
		)

	})

}

//五鬼
func (ui *UI) ly1Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*3),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				ly1 := material.H6(ui.th, ly1Name)
				ly1.Font = text.Font{Typeface: "Noto"}
				ly1.Color = color.NRGBA{0, 0, 0, 255}
				return ly1.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, ly1txt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})

}

//岁煞
func (ui *UI) ly2Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*5),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				ly2 := material.H6(ui.th, ly2Name)
				ly2.Font = text.Font{Typeface: "Noto"}
				ly2.Color = color.NRGBA{0, 0, 0, 255}
				return ly2.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, ly2txt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//岁凶
func (ui *UI) ly3Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*7),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				ly2 := material.H6(ui.th, ly3Name)
				ly2.Font = text.Font{Typeface: "Noto"}
				ly2.Color = color.NRGBA{0, 0, 0, 255}
				return ly2.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, ly3txt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//月表吉
func (ui *UI) lmLayout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*13),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lm := material.H6(ui.th, lmName)
				lm.Font = text.Font{Typeface: "Noto"}
				lm.Color = color.NRGBA{255, 0, 0, 255}
				return lm.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, lmtxt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//月凶
func (ui *UI) lm1Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*16),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lm1 := material.H6(ui.th, lm1Name)
				lm1.Font = text.Font{Typeface: "Noto"}
				lm1.Color = color.NRGBA{0, 0, 0, 255}
				return lm1.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, lm1txt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//日吉
func (ui *UI) ldLayout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*18),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lm1 := material.H6(ui.th, ldName)
				lm1.Font = text.Font{Typeface: "Noto"}
				lm1.Color = color.NRGBA{255, 0, 0, 255}
				return lm1.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, ldtxt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//日凶
func (ui *UI) ld1Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40 + 25*20),
	}
	top := layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				lm1 := material.H6(ui.th, ld1Name)
				lm1.Font = text.Font{Typeface: "Noto"}
				lm1.Color = color.NRGBA{0, 0, 0, 255}
				return lm1.Layout(gtx)
			}),
			top,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, ld1txt)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//24节气
func (ui *UI) jq24Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40),
	}
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(30)}.Layout)
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, zqs)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(ui.th, jqs)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//基本纪年
func (ui *UI) btn3Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40),
	}
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn3s = pub.ArrStringToS(btnArr, "\n")
				body := material.Body1(ui.th, btn3s)
				body.Font = text.Font{Typeface: "Noto"}
				body.Color = color.NRGBA{0, 0, 0, 255}
				return body.Layout(gtx)
			}),
		)
	})
}

//关于
func (ui *UI) btn4Layout(gtx layout.Context) layout.Dimensions {
	in := layout.Inset{
		Top: unit.Dp(40),
	}
	return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		body := material.Body1(ui.th, btn4s)
		body.Font = text.Font{Typeface: "Noto"}
		body.Color = color.NRGBA{0, 0, 0, 255}
		body.TextSize = unit.Dp(16)
		return body.Layout(gtx)
	})
}
func addName() {
	lyName = "岁吉"
	ly1Name = "五鬼"
	ly2Name = "岁煞"
	ly3Name = "岁凶"

	lmName = "月吉"
	lm1Name = "月凶"

	ldName = "日吉"
	ld1Name = "日凶"
}
func cleanjq24() {
	jqs = ""
	zqs = ""
	btn3s = ""
}
func cleanAll() {
	lytxt = ""  //岁吉
	ly1txt = "" //五鬼
	ly2txt = "" //岁煞
	ly3txt = "" //岁凶

	lmtxt = ""  //月吉
	lm1txt = "" //月凶

	ldtxt = ""  //日吉
	ld1txt = "" //日凶
	lyName = ""
	ly1Name = ""
	ly2Name = ""
	ly3Name = ""

	lmName = ""
	lm1Name = ""

	ldName = ""
	ld1Name = ""
	jqs = ""
	zqs = ""
	btnArr = []string{}
	btn3s = ""
	btn4s = ""
}

//
func convF(input string) time.Time {
	ys := input[:4]
	ms := input[4:6]
	ds := input[6:8]
	hs := input[8:10]
	y, _ := strconv.Atoi(ys)
	m, _ := strconv.Atoi(ms)
	d, _ := strconv.Atoi(ds)
	h, _ := strconv.Atoi(hs)
	tx := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	return tx
}

//
func xjbfF(tx time.Time) {
	y, m, d, h := tx.Year(), int(tx.Month()), tx.Day(), tx.Hour()
	obj := gz.NewGanZhi(y, m, d, h)
	ys, ys1, ys2, ys3 := obj.XJBF().NianBiao()
	s1 := pub.ArrStringToS(ys, " ")  //吉
	s2 := pub.ArrStringToS(ys1, " ") //五鬼
	s3 := pub.ArrStringToS(ys2, " ") //煞
	s4 := pub.ArrStringToS(ys3, " ") //凶
	lytxt = s1
	ly1txt = s2
	ly2txt = s3
	ly3txt = s4

	//yzls := obj.XJBF().Yzl() //月总论
	arr5, arr6 := obj.XJBF().YueBiao()
	s5 := pub.ArrStringToS(arr5, " ")
	s6 := pub.ArrStringToS(arr6, " ")
	lmtxt = s5
	lm1txt = s6

	jz60 := gz.GetJzArr()
	arr7, arr8 := obj.XJBF().RiBiao(jz60)
	s7 := pub.ArrStringToS(arr7, " ")
	s8 := pub.ArrStringToS(arr8, " ")
	ldtxt = s7
	ld1txt = s8

}

//24节气
func jq24F(tx time.Time) {
	y, m, d, h := tx.Year(), int(tx.Month()), tx.Day(), tx.Hour()
	obj := gz.NewGanZhi(y, m, d, h)
	jqarr, _ := obj.Jq24()

	var qarr, jarr []string
	for i := 0; i < len(jqarr); i++ {
		if i%2 == 0 {
			qarr = append(qarr, jqarr[i][:18])
		}
		if i%2 == 1 {
			jarr = append(jarr, jqarr[i][:18])
		}
	}
	s9 := pub.ArrStringToS(qarr, "\n")
	jarrs := pub.ArrStringToS(jarr, "\n")
	zqs = s9
	jqs = jarrs
}

//基本纪年
func btn3F(tx time.Time) {
	y, m, d, h := tx.Year(), int(tx.Month()), tx.Day(), tx.Hour()
	weekN := int(tx.Weekday())
	weeks := pub.WeekName(weekN)
	obj := gz.NewGanZhi(y, m, d, h)
	solars := fmt.Sprintf("%d年%d月%d日%d点 %s", y, m, d, h, weeks)
	_, _, _, moons := basic.GetLunar(y, m, d)
	phase := moon.Phase(tx)
	moons = fmt.Sprintf("阴历:%s 月相:%5f", moons, phase)
	gzs := fmt.Sprintf("%s年%s月%s日%s时", obj.YGZ, obj.MGZ, obj.DGZ, obj.HGZ)
	nys := fmt.Sprintf("%s", obj.GetNaYin())
	_, jqsx := obj.Jq24()
	yjzhi, yjName, yjt := obj.YueJiang()
	yjs := fmt.Sprintf("月将:%s(%s): %s", yjzhi, yjName, yjt.Format("2006-01-02"))
	huangHei := obj.RiHuangHei()
	huangHei = fmt.Sprintf("日黄黑:%s", huangHei)
	jianChu := obj.RiJianChu()
	jianChu = fmt.Sprintf("日建除:%s", jianChu)
	riQin := obj.RiQin(weekN)
	riQin = fmt.Sprintf("日禽:%s", riQin)
	huangHeiH := obj.ShiHuangHei()
	huangHeiH = fmt.Sprintf("时黄黑:%s", huangHeiH)
	jqt := obj.Jq24T()
	jueliri := obj.XJBF().JueLiRi(tx, jqt)
	tms, tms1 := obj.GuiDengTianMen() //贵登天门
	if tms != "" || tms1 != "" {
		tms = fmt.Sprintf("贵登天门时:%s %s", tms, tms1)
	}
	zhis, xings, chongs, hais := obj.XJBF().TaiSui()
	dimus := obj.DimuJing()

	if tms != "" {
		btnArr = append(btnArr, solars, gzs, nys, moons, jqsx, yjs, huangHei, jianChu, riQin, huangHeiH, jueliri, tms, zhis, xings, chongs, hais, dimus)

	} else {
		btnArr = append(btnArr, solars, gzs, nys, moons, jqsx, yjs, huangHei, jianChu, riQin, huangHeiH, jueliri, zhis, xings, chongs, hais, dimus)
	}
}

//about
func btn4F() {
	btn4s = `
协纪辩方书E
v0.9.6
参考:协纪辩方书
感谢:JetBrains https://www.jetbrains.com
下载地址 https://github.com/Aquarian-Age/ccal/releases
本程序不含任何担保
`
}
