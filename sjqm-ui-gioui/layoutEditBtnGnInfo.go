/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 *
 */

package main

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func (res *Res) layoutGanZhiS(gtx layout.Context, th *material.Theme) layout.Dimensions {
	inset4 := layout.Inset{Top: unit.Dp(50)}
	return inset4.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gzs := material.H6(th, res.info)
		gzs.TextSize = unit.Sp(14)
		gzs.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
		return gzs.Layout(gtx)
	})
}

func (res *Res) layoutEditAndBtn(th *material.Theme, gtx layout.Context) layout.Dimensions {
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: unit.Dp(26)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					wb := widget.Border{Width: unit.Dp(0.6)}
					wb.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &res.ed1, "year").Layout(gtx)
					})
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					wb := widget.Border{Width: unit.Dp(0.6)}
					wb.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &res.ed2, "month").Layout(gtx)
					})
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					wb := widget.Border{Width: unit.Dp(0.6)}
					wb.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &res.ed3, "day").Layout(gtx)
					})
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					wb := widget.Border{Width: unit.Dp(0.6)}
					wb.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &res.ed4, "hour").Layout(gtx)
					})
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					wb := widget.Border{Width: unit.Dp(0.6)}
					wb.Color = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &res.ed5, "min").Layout(gtx)
					})
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btnl := material.Button(th, &res.btn1, "JQH")
					btnl.Inset.Top = unit.Dp(3)
					btnl.Inset.Bottom = unit.Dp(1)
					btnl.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return btnl.Layout(gtx)
				}),
				space,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btnam := material.Button(th, &res.btna, "About")
					btnam.Inset.Top = unit.Dp(3)
					btnam.Inset.Bottom = unit.Dp(1)
					btnam.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
					return btnam.Layout(gtx)
				}),
			)
		}),
	)
}

func (res *Res) layoutGnInfo(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: top1 + 90}
	spaceg := layout.Rigid(layout.Spacer{Width: unit.Dp(3)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng1 := material.Button(th, &res.bg1, "坎")
						btng1.Inset.Top = unit.Dp(3)
						btng1.Inset.Bottom = unit.Dp(1)
						btng1.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng1.Layout(gtx)
					}),
					spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng8 := material.Button(th, &res.bg8, "艮")
						btng8.Inset.Top = unit.Dp(3)
						btng8.Inset.Bottom = unit.Dp(1)
						btng8.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng8.Layout(gtx)
					}),
					spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng3 := material.Button(th, &res.bg3, "震")
						btng3.Inset.Top = unit.Dp(3)
						btng3.Inset.Bottom = unit.Dp(1)
						btng3.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng3.Layout(gtx)
					}),
					spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng4 := material.Button(th, &res.bg4, "巽")
						btng4.Inset.Top = unit.Dp(3)
						btng4.Inset.Bottom = unit.Dp(1)
						btng4.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng4.Layout(gtx)
					}),
					spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng9 := material.Button(th, &res.bg9, "离")
						btng9.Inset.Top = unit.Dp(3)
						btng9.Inset.Bottom = unit.Dp(1)
						btng9.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng9.Layout(gtx)
					}),
					spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng2 := material.Button(th, &res.bg2, "坤")
						btng2.Inset.Top = unit.Dp(3)
						btng2.Inset.Bottom = unit.Dp(1)
						btng2.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng2.Layout(gtx)
					}), spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng7 := material.Button(th, &res.bg7, "兑")
						btng7.Inset.Top = unit.Dp(3)
						btng7.Inset.Bottom = unit.Dp(1)
						btng7.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng7.Layout(gtx)
					}), spaceg,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btng6 := material.Button(th, &res.bg6, "乾")
						btng6.Inset.Top = unit.Dp(3)
						btng6.Inset.Bottom = unit.Dp(1)
						btng6.Background = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
						return btng6.Layout(gtx)
					}),
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Editor(th, ed, "")
				body.TextSize = unit.Sp(14)
				ed.SetText(res.btnGxs)
				return body.Layout(gtx)
			}),
		)
	})
}

var ed = new(widget.Editor)
