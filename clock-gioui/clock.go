package main

import (
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(480), unit.Dp(480)),
			app.Title("clock"),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				//------------背景颜色
				bgCol := color.NRGBA{R: 0, G: 191, B: 255, A: 255}
				paint.ColorOp{Color: bgCol}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				//------------

				l := material.H1(th, time.Now().Local().Format("15:04:05"))
				maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
				l.Color = maroon
				l.Alignment = text.Middle
				l.Layout(gtx)

				e.Frame(gtx.Ops)
			}
		case <-ticker.C:
			w.Invalidate()
		}
	}
}
