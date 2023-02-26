/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 *
 */

package main

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

var (
	top    = unit.Dp(150)
	top3   = unit.Dp(245)
	left5  = unit.Dp(110)
	left2  = unit.Dp(205)
	top1   = unit.Dp(340)
	right2 = unit.Dp(6)   //适配移动端App
	bottom = unit.Dp(120) //120适配移动端App

	DarkOrchid = color.NRGBA{R: 127, G: 0, B: 255, A: 255}
	wb         = widget.Border{Width: unit.Dp(0.6), Color: DarkOrchid}
)

//矩形
func strokeRect(gtx layout.Context) {
	const r = 10
	rect := clip.RRect{
		Rect: image.Rectangle{
			Min: image.Point{X: gtx.Dp(unit.Dp(1)), Y: gtx.Dp(unit.Dp(150))}, //动态适配
			Max: image.Point{X: gtx.Dp(unit.Dp(310)), Y: gtx.Dp(unit.Dp(430))},
		},
		SE: r, SW: r, NW: r, NE: r}
	ops := gtx.Ops
	paint.FillShape(ops,
		DarkOrchid,
		clip.Stroke{
			Path:  rect.Path(ops),
			Width: 2,
		}.Op(),
	)
}

//横线
func strokeTriangle(gtx layout.Context) {
	var path clip.Path
	ops := gtx.Ops
	path.Begin(ops)
	path.MoveTo(f32.Pt(float32(gtx.Dp(unit.Dp(1))), float32(gtx.Dp(unit.Dp(150+95)))))
	path.LineTo(f32.Pt(float32(gtx.Dp(unit.Dp(310))), float32(gtx.Dp(unit.Dp(150+95)))))
	path.Close()
	paint.FillShape(ops,
		DarkOrchid,
		clip.Stroke{
			Path:  path.End(),
			Width: 1,
		}.Op())
}
func strokeTriangle2(gtx layout.Context) {
	var path clip.Path
	ops := gtx.Ops
	path.Begin(ops)
	path.MoveTo(f32.Pt(float32(gtx.Dp(unit.Dp(1))), float32(gtx.Dp(unit.Dp(150+95*2)))))
	path.LineTo(f32.Pt(float32(gtx.Dp(unit.Dp(310))), float32(gtx.Dp(unit.Dp(150+95*2)))))
	path.Close()
	paint.FillShape(ops,
		DarkOrchid,
		clip.Stroke{
			Path:  path.End(),
			Width: 1,
		}.Op())
}

//竖线
func strokeTriangle3(gtx layout.Context) {
	var path clip.Path
	ops := gtx.Ops
	path.Begin(ops)
	path.MoveTo(f32.Pt(float32(gtx.Dp(unit.Dp(108))), float32(gtx.Dp(unit.Dp(150)))))
	path.LineTo(f32.Pt(float32(gtx.Dp(unit.Dp(108))), float32(gtx.Dp(unit.Dp(450)))))
	path.Close()
	paint.FillShape(ops,
		DarkOrchid,
		clip.Stroke{
			Path:  path.End(),
			Width: 1,
		}.Op())
}
func strokeTriangle4(gtx layout.Context) {
	var path clip.Path
	ops := gtx.Ops
	path.Begin(ops)
	path.MoveTo(f32.Pt(float32(gtx.Dp(unit.Dp(203))), float32(gtx.Dp(unit.Dp(150)))))
	path.LineTo(f32.Pt(float32(gtx.Dp(unit.Dp(203))), float32(gtx.Dp(unit.Dp(450)))))
	path.Close()
	paint.FillShape(ops,
		DarkOrchid,
		clip.Stroke{
			Path:  path.End(),
			Width: 1,
		}.Op())
}

//4宫
func (res *Res) show4(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset4 := layout.Inset{Top: top, Left: unit.Dp(3), Right: unit.Dp(right2), Bottom: unit.Dp(bottom)}
	return inset4.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God4 := material.H6(th, res.baShen[3])
				God4.TextSize = unit.Sp(14)
				return God4.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star := material.H6(th, res.jiuXing[3]+" "+res.tianQi[3])
				star.TextSize = unit.Sp(14)
				return star.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door4 := material.H6(th, res.baMen[3]+" "+res.diQi[3])
				door4.TextSize = unit.Sp(14)
				return door4.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz4 := material.H6(th, res.anGan[3])
				gz4.TextSize = unit.Sp(14)
				return gz4.Layout(gtx)
			}),
		)
	})
}

//9宫
func (res *Res) show9(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset9 := layout.Inset{Top: top, Left: left5, Right: unit.Dp(right2), Bottom: unit.Dp(bottom)}
	return inset9.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God9 := material.H6(th, res.baShen[4])
				God9.TextSize = unit.Sp(14)
				return God9.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star9 := material.H6(th, res.jiuXing[4]+" "+res.tianQi[4])
				star9.TextSize = unit.Sp(14)
				return star9.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door9 := material.H6(th, res.baMen[4]+" "+res.diQi[4])
				door9.TextSize = unit.Sp(14)
				return door9.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz9 := material.H6(th, res.anGan[4])
				gz9.TextSize = unit.Sp(14)
				return gz9.Layout(gtx)
			}),
		)
	})
}

//2宫
func (res *Res) show2(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset2 := layout.Inset{Top: top, Left: left2, Right: unit.Dp(right2), Bottom: unit.Dp(bottom)}
	return inset2.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God2 := material.H6(th, res.baShen[5])
				God2.TextSize = unit.Sp(14)
				return God2.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star2 := material.H6(th, res.jiuXing[5]+" "+res.tianQi[5])
				star2.TextSize = unit.Sp(14)
				return star2.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door2 := material.H6(th, res.baMen[5]+" "+res.diQi[5])
				door2.TextSize = unit.Sp(14)
				return door2.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz2 := material.H6(th, res.anGan[5])
				gz2.TextSize = unit.Sp(14)
				return gz2.Layout(gtx)
			}),
		)
	})
}

//3宫
func (res *Res) show3(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset3 := layout.Inset{Top: top3, Left: unit.Dp(3), Right: unit.Dp(right2), Bottom: unit.Dp(bottom)}
	return inset3.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God3 := material.H6(th, res.baShen[2])
				God3.TextSize = unit.Sp(14)
				return God3.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star3 := material.H6(th, res.jiuXing[2]+" "+res.tianQi[2])
				star3.TextSize = unit.Sp(14)
				return star3.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door3 := material.H6(th, res.baMen[2]+" "+res.diQi[2])
				door3.TextSize = unit.Sp(14)
				return door3.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz3 := material.H6(th, res.anGan[2])
				gz3.TextSize = unit.Sp(14)
				return gz3.Layout(gtx)
			}),
		)
	})
}

//5宫
func (res *Res) show5(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset5 := layout.Inset{Top: top3, Left: left5}
	return inset5.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star5 := material.H6(th, res.g5s)
				star5.TextSize = unit.Sp(14)
				return star5.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e5 := material.H6(th, res.diQi[8])
				e5.TextSize = unit.Sp(14)
				return e5.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz5 := material.H6(th, res.anGan[8])
				gz5.TextSize = unit.Sp(14)
				return gz5.Layout(gtx)
			}),
		)
	})
}

//7宫
func (res *Res) show7(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset7 := layout.Inset{Top: top3, Left: left2}
	return inset7.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God7 := material.H6(th, res.baShen[6])
				God7.TextSize = unit.Sp(14)
				return God7.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star7 := material.H6(th, res.jiuXing[6]+" "+res.tianQi[6])
				star7.TextSize = unit.Sp(14)
				return star7.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door7 := material.H6(th, res.baMen[6]+" "+res.diQi[6])
				door7.TextSize = unit.Sp(14)
				return door7.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz7 := material.H6(th, res.anGan[6])
				gz7.TextSize = unit.Sp(14)
				return gz7.Layout(gtx)
			}),
		)
	})
}

//8宫
func (res *Res) show8(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset8 := layout.Inset{Top: top1, Left: unit.Dp(3), Right: unit.Dp(right2), Bottom: unit.Dp(bottom)}
	return inset8.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God8 := material.H6(th, res.baShen[1])
				God8.TextSize = unit.Sp(14)
				return God8.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star8 := material.H6(th, res.jiuXing[1]+" "+res.tianQi[1])
				star8.TextSize = unit.Sp(14)
				return star8.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door8 := material.H6(th, res.baMen[1]+" "+res.diQi[1])
				door8.TextSize = unit.Sp(14)
				return door8.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz8 := material.H6(th, res.anGan[1])
				gz8.TextSize = unit.Sp(14)
				return gz8.Layout(gtx)
			}),
		)
	})
}

//1宫
func (res *Res) show1(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset1 := layout.Inset{Top: top1, Left: left5}
	return inset1.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God1 := material.H6(th, res.baShen[0])
				God1.TextSize = unit.Sp(14)
				return God1.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star1 := material.H6(th, res.jiuXing[0]+" "+res.tianQi[0])
				star1.TextSize = unit.Sp(14)
				return star1.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door1 := material.H6(th, res.baMen[0]+" "+res.diQi[0])
				door1.TextSize = unit.Sp(14)
				return door1.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz1 := material.H6(th, res.anGan[0])
				gz1.TextSize = unit.Sp(14)
				return gz1.Layout(gtx)
			}),
		)
	})
}

//6宫
func (res *Res) show6(th *material.Theme, gtx layout.Context) layout.Dimensions {
	inset6 := layout.Inset{Top: top1, Left: left2}
	return inset6.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				God6 := material.H6(th, res.baShen[7])
				God6.TextSize = unit.Sp(14)
				return God6.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				star6 := material.H6(th, res.jiuXing[7]+" "+res.tianQi[7])
				star6.TextSize = unit.Sp(14)
				return star6.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				door6 := material.H6(th, res.baMen[7]+" "+res.diQi[7])
				door6.TextSize = unit.Sp(14)
				return door6.Layout(gtx)
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gz6 := material.H6(th, res.anGan[7])
				gz6.TextSize = unit.Sp(14)
				return gz6.Layout(gtx)
			}),
		)
	})
}
