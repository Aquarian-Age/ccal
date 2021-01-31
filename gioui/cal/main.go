package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gonoto/notosans"
	"liangzi.local/cal/cal"
)

func main() {
	ui := NewUI()
	go func() {
		w := app.NewWindow(app.Title("干支历"), app.Size(unit.Dp(500), unit.Dp(600)))
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	Theme *material.Theme
	edits Edit
}
type Edit struct {
	ed1, ed2, ed3, ed4 widget.Editor
	btn                widget.Clickable
	btn2               widget.Clickable
	btn3               widget.Clickable
}

func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			ui.LayoutUI(gtx)
			e.Frame(gtx.Ops)
		case key.Event:
			switch e.Name {
			case key.NameEscape:
				return nil
			}
		case system.DestroyEvent:
			return e.Err
		}
	}
	return nil
}

func NewUI() *UI {
	ui := &UI{}
	ui.Theme = utf8Font()
	ui.edits.Inits()
	return ui
}

var s = ymd(time.Now().Local())

func (ui *UI) LayoutUI(gtx layout.Context) layout.Dimensions {
	inset := layout.UniformInset(unit.Dp(12))
	inset.Top = unit.Dp(26) //错开Android顶部栏
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.edits.LayoutEdit(gtx, ui.Theme)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return ui.edits.LayoutShow(gtx, ui.Theme)
			}),
		)
	})
}

///
func (ed *Edit) LayoutShow(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if ed.btn.Clicked() {
		y := convEdit(ed.ed1.Text())
		m := convEdit(ed.ed2.Text())
		d := convEdit(ed.ed3.Text())
		h := convEdit(ed.ed4.Text())
		t := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		eds := ymd(t)
		s = eds
	} else if ed.btn2.Clicked() {
		y := convEdit(ed.ed1.Text())
		m := convEdit(ed.ed2.Text())
		d := convEdit(ed.ed3.Text())
		eds := yueli(y, m, d)
		s = eds
	} else if ed.btn3.Clicked() {
		y := convEdit(ed.ed1.Text())
		eds := jq24(y)
		s = eds
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			body := material.Body1(th, s)
			body.Font = text.Font{Typeface: "Noto"}
			body.TextSize = unit.Dp(14)
			body.Alignment = text.Start
			return body.Layout(gtx)
		}),
	)
}

///
func (ed *Edit) LayoutEdit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	borderWidth := float32(0.5)
	borderColor := color.NRGBA{A: 107}
	switch {
	case ed.ed1.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed2.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed3.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.ed4.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	}
	wb := widget.Border{
		Color:        borderColor,
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(borderWidth),
	}
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout) //小部件间距
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed1, "year").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed2, "month").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed3, "day").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.ed4, "hour").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &ed.btn, "纪年")
			btn.Font = text.Font{Typeface: "Noto"}
			return btn.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn2 := material.Button(th, &ed.btn2, "月历")
			btn2.Font = text.Font{Typeface: "Noto"}
			return btn2.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn3 := material.Button(th, &ed.btn3, "24节气")
			btn3.Font = text.Font{Typeface: "Noto"}
			return btn3.Layout(gtx)
		}),
	)
}
func (ed *Edit) Inits() {
	ed.ed1.SingleLine = true
	ed.ed2.SingleLine = true
	ed.ed3.SingleLine = true
	ed.ed4.SingleLine = true
}

//
func utf8Font() *material.Theme {
	fonts := gofont.Collection()
	fonts = appendOTC(fonts, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(fonts)
	return th
}

func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

///农历年月日时 干支年月日时 当前节气 朔望上弦下弦
func ymd(t time.Time) string {
	lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
	var lunars string
	if lunar.LeapSday == 0 {
		lunars = fmt.Sprintf("%d年%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LM), lunar.Ydx, lunar.LRmc, lunar.Week)
	} else {
		lunars = fmt.Sprintf("%d年闰%s月(%s)%s %s",
			lunar.LY, cal.AliasLunarMonth(lunar.LeapM), lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
	}
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	gzs := gz.YearGZ + "年 " + gz.MonthGZ + "月 " + gz.DayGZ + "日 " + gz.HourGZ + "时"

	arrt := cal.YueJiangJQArrT(t.Year())
	jq := cal.NewYueJiangJQ(arrt)
	todayJQs := jq.TodayJQ(t.Year(), t)

	sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
	shuo := "朔: " + sw.ShuoTS
	wang := "望: " + sw.WangTS
	shang := "上弦: " + sw.ShangXianTS
	xia := "下弦: " + sw.XiaXianTS
	s1 := lunars + "\n" + gzs + "\n"
	s2 := shuo + "\n" + wang + "\n" + shang + "\n" + xia + "\n"
	s3 := "节气: " + todayJQs

	dmjs := cal.Dmj(t.Year(), int(t.Month()), t.Day())
	s := "\n" + s1 + s2 + s3 + "\n\n" + dmjs

	return s + "\n" + "UI地址: https://github.com/Aquarian-Age/ccal"
}

//月历
func yueli(y, m, d int) string {
	yl := cal.NewYueLi(y, m, d)
	l := yl.LunarD
	s := yl.SolarD
	gz := yl.GzD
	var s1, l1, g1, s2, l2, g2, s3, l3, g3, s4, l4, g4, s5, l5, g5 []string
	for i := 0; i < 6; i++ {
		s1 = append(s1, s[i])
		l1 = append(l1, l[i])
		g1 = append(g1, gz[i])
	}
	for i := 6; i < 12; i++ {
		s2 = append(s2, s[i])
		l2 = append(l2, l[i])
		g2 = append(g2, gz[i])
	}
	for i := 12; i < 18; i++ {
		s3 = append(s3, s[i])
		l3 = append(l3, l[i])
		g3 = append(g3, gz[i])
	}
	for i := 18; i < 24; i++ {
		s4 = append(s4, s[i])
		l4 = append(l4, l[i])
		g4 = append(g4, gz[i])
	}
	for i := 24; i < len(l); i++ {
		s5 = append(s5, s[i])
		l5 = append(l5, l[i])
		g5 = append(g5, gz[i])

	}

	rs1 := strings.Join(s1, " | ")
	rl1 := strings.Join(l1, " | ")
	rg1 := strings.Join(g1, " | ")
	rs2 := strings.Join(s2, " | ")
	rl2 := strings.Join(l2, " | ")
	rg2 := strings.Join(g2, " | ")
	rs3 := strings.Join(s3, " | ")
	rl3 := strings.Join(l3, " | ")
	rg3 := strings.Join(g3, " | ")
	rs4 := strings.Join(s4, " | ")
	rl4 := strings.Join(l4, " | ")
	rg4 := strings.Join(g4, " | ")
	rs5 := strings.Join(s5, " | ")
	rl5 := strings.Join(l5, " | ")
	rg5 := strings.Join(g5, " | ")

	rs := "\n" + "---------------------------------------------------------" + "\n" +
		rs1 + "\n" + rl1 + "\n" + rg1 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs2 + "\n" + rl2 + "\n" + rg2 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs3 + "\n" + rl3 + "\n" + rg3 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs4 + "\n" + rl4 + "\n" + rg4 +
		"\n" + "---------------------------------------------------------" + "\n" +
		rs5 + "\n" + rl5 + "\n" + rg5

	return rs + "\n" + "UI地址: https://github.com/Aquarian-Age/ccal"
}

//24节气
func jq24(y int) string {
	jqt := cal.YueJiangJQArrT(y)
	yjq := cal.NewYueJiangJQ(jqt)
	return yjq.JQ24()
}
func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}
