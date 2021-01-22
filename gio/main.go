package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"gioui.org/font/opentype"
	"liangzi.local/cal/cal"
)

func init() {

}

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := mth()
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			t := time.Now().Local()
			lunar := cal.NewLunarMonth(t.Year(), int(t.Month()), t.Day())
			var info string
			if lunar.LeapSday == 0 {
				info = fmt.Sprintf("%d年-%d月(%s)-%s %s",
					lunar.LY, lunar.LM, lunar.Ydx, lunar.LRmc, lunar.Week)
			} else {
				info = fmt.Sprintf("%d年-闰%d月(%s)-%s %s",
					lunar.LY, lunar.LeapM, lunar.LeapYdx, lunar.LeapRmc, lunar.Week)
			}
			gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
			gzs := gz.YearGZ + "-" + gz.MonthGZ + "-" + gz.DayGZ + "-" + gz.HourGZ

			sw := cal.NewShuoWangTS(t.Year(), int(t.Month()), t.Day())
			shuo := "朔: " + sw.ShuoTS
			wang := "望: " + sw.WangTS
			shang := "上弦: " + sw.ShangXianTS
			xia := "下弦: " + sw.XiaXianTS
			sws := shuo + "\n" + wang + "\n" + shang + "\n" + xia
			s := info + "\n" + gzs + "\n" + sws

			//l := material.H1(th, info)
			l := material.H4(th, s)
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			l.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

//返回一个中文字体主题
func mth() *material.Theme {
	f, err := os.Open("/opt/fonts/yahei.ttf")
	if err != nil {
		log.Fatal(err)
	}
	//返回(opentype*Collection, error)
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}

	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}
