package main

import (
	"image"
	"strings"

	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type wly func(gtx layout.Context) layout.Dimensions

func (view wly) Run(w *Window) error {
	var ops op.Ops

	applicationClose := w.App.Context.Done()
	for {
		select {
		case <-applicationClose:
			return nil
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				view(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}
}
func (v *Letters) LayoutN(gtx layout.Context) layout.Dimensions {
	hGrid := outlay.Grid{
		Num:  7, //四方七宿
		Axis: layout.Horizontal,
	}
	th := v.win.App.Theme

	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, len(v.items), func(gtx layout.Context, index int) layout.Dimensions {
				item := v.items[index]

				if item.Click.Clicked() {
					for i := 0; i < len(starNames); i++ {
						if strings.EqualFold(starNames[i], item.Text) {
							bfyqRQ(item.Text)
							break
						}
					}
				}
				///
				var nameWidth int
				return layout.Stack{Alignment: layout.S}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &item.Click, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(6)).Layout(gtx,
								material.Body1(th, item.Text).Layout,
							)
						})
						nameWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						nameHeight := gtx.Px(unit.Dp(2))
						return layout.Dimensions{
							Size: image.Point{X: nameWidth, Y: nameHeight},
						}
					}),
				)
			})
		}),
	)
}
