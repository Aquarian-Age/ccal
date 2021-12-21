/*
 * 版权所有 2021.
 * 本人保留所有权。
 *  本文档及其相关产品的使用、复制、分发和反编译均受许可证限制。
 *  未经作者事先书面许可,不得以任何形式、任何手段复制本产品或文档的任何部分。
 *
 */

package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"github.com/gonoto/notosans"
	"liangzi.local/qx/pkg/gz"
	"liangzi.local/qx/pkg/pub"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type UI struct {
	edit1, edit2, edit3, edit4 widget.Editor
	btna                       widget.Clickable
	btnb                       widget.Clickable
	theme                      *material.Theme
	lables                     []material.LabelStyle
}

var (
	t    = time.Now().Local()
	ymds string
	y    = t.Year()
	m    = int(t.Month())
	d    = t.Day()
)

func main() {
	ui := NewUI()
	go func() {
		w := app.NewWindow(
			app.Title("star"),
			app.Size(unit.Dp(360), unit.Dp(640)), //360x640 asus_ze552kl
		)
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

func NewUI() *UI {
	ui := &UI{}
	ui.theme = utf8Font() // material.NewTheme(gofont.Collection())
	ui.Inits()
	return ui
}
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
func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			ui.showTime(gtx)
			//
			ui.showStars(gtx)
			ui.showStar1(gtx)
			ui.showStar2(gtx)
			ui.showStar3(gtx)
			switch day {
			case 29:
				ui.showStar29(gtx)
			case 30:
				ui.showStar30(gtx)
			case 31:
				ui.showStar31(gtx)
			}

			ui.showInfo(gtx)
			//------------------
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

////
func (ui *UI) showInfo(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(480)}
	return inset.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					l := material.H6(ui.theme, ymds)
					l.Font = text.Font{Typeface: "Noto"}
					l.TextSize = unit.Dp(18)
					maroon := color.NRGBA{R: 255, G: 69, B: 0, A: 255}
					l.Color = maroon
					l.Alignment = text.Middle
					return l.Layout(gtx)
				}),
			)
		},
	)
}

//
func (ui *UI) showTime(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		//实时显示时间
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.H6(ui.theme, time.Now().Local().Format("2006-01-02 15:04:05"))
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			return l.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.LayoutEdit(gtx, ui.theme)
		}),
	)
}
func (ed *UI) Inits() {
	ed.edit1.SingleLine = true
	ed.edit2.SingleLine = true
	ed.edit3.SingleLine = true
	ed.edit4.SingleLine = true
}

//
func (ed *UI) LayoutEdit(gtx layout.Context, th *material.Theme) layout.Dimensions {
	borderWidth := float32(0.5)
	borderColor := color.NRGBA{A: 107}
	switch {
	case ed.edit1.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.edit2.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.edit3.Focused():
		borderColor = th.Palette.ContrastBg
		borderWidth = 1
	case ed.edit4.Focused():
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
					material.Editor(th, &ed.edit1, "year").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.edit2, "month").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.edit3, "day").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wb.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(9)).Layout(gtx,
					material.Editor(th, &ed.edit4, "hour").Layout)
			})
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th, &ed.btna, "OK")
			return btn.Layout(gtx)
		}),
		space,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btn2 := material.Button(th, &ed.btnb, "about")
			return btn2.Layout(gtx)
		}),
	)
}

//
func (ed *UI) LayoutShow(gtx layout.Context, th *material.Theme) layout.Dimensions {
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			l := material.H6(th, ymds)
			l.Font = text.Font{Typeface: "Noto"}
			l.TextSize = unit.Dp(18)
			maroon := color.NRGBA{R: 255, G: 69, B: 0, A: 255}
			l.Color = maroon
			l.Alignment = text.Middle
			return l.Layout(gtx)
		}),
		space,
	)
}

func (ui *UI) showStars(gtx layout.Context) layout.Dimensions {
	if ui.btna.Clicked() {
		if ui.edit1.Text() == "" || ui.edit2.Text() == "" || ui.edit3.Text() == "" || ui.edit4.Text() == "" {
			t = time.Now().Local()
			y, m, d = t.Year(), int(t.Month()), t.Day()
		} else {
			y = convEdit(ui.edit1.Text())
			m = convEdit(ui.edit2.Text())
			d = convEdit(ui.edit3.Text())
			h := convEdit(ui.edit4.Text())
			t = time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
		}
		ymds = t.String()
		ui.stars(t)
	} else {
		ui.stars(t)
	}
	if ui.btnb.Clicked() {
		ymds = `下载地址:
		https://github.com/Aquarian-Age/ccal`

	}
	inset := layout.Inset{Top: unit.Dp(80)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[0].Font = text.Font{Typeface: "Noto"}
				solars[0].Text = arr[0] + weeks[0] + "\n" + dgzarr[0] + "\n" + jcarr[0] + hharr[0] + "\n" + riQinarr[0]
				solars[0].TextSize = size
				return solars[0].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[1].Font = text.Font{Typeface: "Noto"}
				solars[1].Text = arr[1] + weeks[1] + "\n" + dgzarr[1] + "\n" + jcarr[1] + hharr[1] + "\n" + riQinarr[1]
				solars[1].TextSize = size
				return solars[1].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[2].Font = text.Font{Typeface: "Noto"}
				solars[2].Text = arr[2] + weeks[2] + "\n" + dgzarr[2] + "\n" + jcarr[2] + hharr[2] + "\n" + riQinarr[2]
				solars[2].TextSize = size
				return solars[2].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[3].Font = text.Font{Typeface: "Noto"}
				solars[3].Text = arr[3] + weeks[3] + "\n" + dgzarr[3] + "\n" + jcarr[3] + hharr[3] + "\n" + riQinarr[3]
				solars[3].TextSize = size
				return solars[3].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[4].Font = text.Font{Typeface: "Noto"}
				solars[4].Text = arr[4] + weeks[4] + "\n" + dgzarr[4] + "\n" + jcarr[4] + hharr[4] + "\n" + riQinarr[4]
				solars[4].TextSize = size
				return solars[4].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[5].Font = text.Font{Typeface: "Noto"}
				solars[5].Text = arr[5] + weeks[5] + "\n" + dgzarr[5] + "\n" + jcarr[5] + hharr[5] + "\n" + riQinarr[5]
				solars[5].TextSize = size
				return solars[5].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[6].Font = text.Font{Typeface: "Noto"}
				solars[6].Text = arr[6] + weeks[6] + "\n" + dgzarr[6] + "\n" + jcarr[6] + hharr[6] + "\n" + riQinarr[6]
				solars[6].TextSize = size
				return solars[6].Layout(gtx)
			}),
			space,
		)
	})

}
func (ui *UI) showStar1(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(160)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[7].Font = text.Font{Typeface: "Noto"}
				solars[7].Text = arr[7] + weeks[7] + "\n" + dgzarr[7] + "\n" + jcarr[7] + hharr[7] + "\n" + riQinarr[7]
				solars[7].TextSize = size
				return solars[7].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[8].Font = text.Font{Typeface: "Noto"}
				solars[8].Text = arr[8] + weeks[8] + "\n" + dgzarr[8] + "\n" + jcarr[8] + hharr[8] + "\n" + riQinarr[8]
				solars[8].TextSize = size
				return solars[8].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[9].Font = text.Font{Typeface: "Noto"}
				solars[9].Text = arr[9] + weeks[9] + "\n" + dgzarr[9] + "\n" + jcarr[9] + hharr[9] + "\n" + riQinarr[9]
				solars[9].TextSize = size
				return solars[9].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[10].Font = text.Font{Typeface: "Noto"}
				solars[10].Text = arr[10] + weeks[10] + "\n" + dgzarr[10] + "\n" + jcarr[10] + hharr[10] + "\n" + riQinarr[10]
				solars[10].TextSize = size
				return solars[10].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[11].Font = text.Font{Typeface: "Noto"}
				solars[11].Text = arr[11] + weeks[11] + "\n" + dgzarr[11] + "\n" + jcarr[11] + hharr[11] + "\n" + riQinarr[11]
				solars[11].TextSize = size
				return solars[11].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[12].Font = text.Font{Typeface: "Noto"}
				solars[12].Text = arr[12] + weeks[12] + "\n" + dgzarr[12] + "\n" + jcarr[12] + hharr[12] + "\n" + riQinarr[12]
				solars[12].TextSize = size
				return solars[12].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[13].Font = text.Font{Typeface: "Noto"}
				solars[13].Text = arr[13] + weeks[13] + "\n" + dgzarr[13] + "\n" + jcarr[13] + hharr[13] + "\n" + riQinarr[13]
				solars[13].TextSize = size
				return solars[13].Layout(gtx)
			}),
			space,
		)
	})
}
func (ui *UI) showStar2(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(240)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[14].Font = text.Font{Typeface: "Noto"}
				solars[14].Text = arr[14] + weeks[14] + "\n" + dgzarr[14] + "\n" + jcarr[14] + hharr[14] + "\n" + riQinarr[14]
				solars[14].TextSize = size
				return solars[14].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[15].Font = text.Font{Typeface: "Noto"}
				solars[15].Text = arr[15] + weeks[15] + "\n" + dgzarr[15] + "\n" + jcarr[15] + hharr[15] + "\n" + riQinarr[15]
				solars[15].TextSize = size
				return solars[15].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[16].Font = text.Font{Typeface: "Noto"}
				solars[16].Text = arr[16] + weeks[16] + "\n" + dgzarr[16] + "\n" + jcarr[16] + hharr[16] + "\n" + riQinarr[16]
				solars[16].TextSize = size
				return solars[16].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[17].Font = text.Font{Typeface: "Noto"}
				solars[17].Text = arr[17] + weeks[17] + "\n" + dgzarr[17] + "\n" + jcarr[17] + hharr[17] + "\n" + riQinarr[17]
				solars[17].TextSize = size
				return solars[17].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[18].Font = text.Font{Typeface: "Noto"}
				solars[18].Text = arr[18] + weeks[18] + "\n" + dgzarr[18] + "\n" + jcarr[18] + hharr[18] + "\n" + riQinarr[18]
				solars[18].TextSize = size
				return solars[18].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[19].Font = text.Font{Typeface: "Noto"}
				solars[19].Text = arr[19] + weeks[19] + "\n" + dgzarr[19] + "\n" + jcarr[19] + hharr[19] + "\n" + riQinarr[19]
				solars[19].TextSize = size
				return solars[19].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[20].Font = text.Font{Typeface: "Noto"}
				solars[20].Text = arr[20] + weeks[20] + "\n" + dgzarr[20] + "\n" + jcarr[20] + hharr[20] + "\n" + riQinarr[20]
				solars[20].TextSize = size
				return solars[20].Layout(gtx)
			}),
			space,
		)
	})
}
func (ui *UI) showStar3(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(320)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[21].Font = text.Font{Typeface: "Noto"}
				solars[21].Text = arr[21] + weeks[21] + "\n" + dgzarr[21] + "\n" + jcarr[21] + hharr[21] + "\n" + riQinarr[21]
				solars[21].TextSize = size
				return solars[21].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[22].Font = text.Font{Typeface: "Noto"}
				solars[22].Text = arr[22] + weeks[22] + "\n" + dgzarr[22] + "\n" + jcarr[22] + hharr[22] + "\n" + riQinarr[22]
				solars[22].TextSize = size
				return solars[22].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[23].Font = text.Font{Typeface: "Noto"}
				solars[23].Text = arr[23] + weeks[23] + "\n" + dgzarr[23] + "\n" + jcarr[23] + hharr[23] + "\n" + riQinarr[23]
				solars[23].TextSize = size
				return solars[23].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[24].Font = text.Font{Typeface: "Noto"}
				solars[24].Text = arr[24] + weeks[24] + "\n" + dgzarr[24] + "\n" + jcarr[24] + hharr[24] + "\n" + riQinarr[24]
				solars[24].TextSize = size
				return solars[24].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[25].Font = text.Font{Typeface: "Noto"}
				solars[25].Text = arr[25] + weeks[25] + "\n" + dgzarr[25] + "\n" + jcarr[25] + hharr[25] + "\n" + riQinarr[25]
				solars[25].TextSize = size
				return solars[25].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[26].Font = text.Font{Typeface: "Noto"}
				solars[26].Text = arr[26] + weeks[26] + "\n" + dgzarr[26] + "\n" + jcarr[26] + hharr[26] + "\n" + riQinarr[26]
				solars[26].TextSize = size
				return solars[26].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[27].Font = text.Font{Typeface: "Noto"}
				solars[27].Text = arr[27] + weeks[27] + "\n" + dgzarr[27] + "\n" + jcarr[27] + hharr[27] + "\n" + riQinarr[27]
				solars[27].TextSize = size
				return solars[27].Layout(gtx)
			}),
			space,
		)
	})
}
func (ui *UI) showStar29(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(400)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[28].Font = text.Font{Typeface: "Noto"}
				solars[28].Text = arr[28] + weeks[28] + "\n" + dgzarr[28] + "\n" + jcarr[28] + hharr[28] + "\n" + riQinarr[28]
				solars[28].TextSize = size
				return solars[28].Layout(gtx)
			}),
			space,
		)
	})
}
func (ui *UI) showStar30(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(400)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[28].Font = text.Font{Typeface: "Noto"}
				solars[28].Text = arr[28] + weeks[28] + "\n" + dgzarr[28] + "\n" + jcarr[28] + hharr[28] + "\n" + riQinarr[28]
				solars[28].TextSize = size
				return solars[28].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[29].Font = text.Font{Typeface: "Noto"}
				solars[29].Text = arr[29] + weeks[29] + "\n" + dgzarr[29] + "\n" + jcarr[29] + hharr[29] + "\n" + riQinarr[29]
				solars[29].TextSize = size
				return solars[29].Layout(gtx)
			}),
			space,
		)
	})
}
func (ui *UI) showStar31(gtx layout.Context) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(400)}
	size := unit.Dp(12)
	space := layout.Rigid(layout.Spacer{Width: unit.Dp(13)}.Layout)
	return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[28].Font = text.Font{Typeface: "Noto"}
				solars[28].Text = arr[28] + weeks[28] + "\n" + dgzarr[28] + "\n" + jcarr[28] + hharr[28] + "\n" + riQinarr[28]
				solars[28].TextSize = size
				return solars[28].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[29].Font = text.Font{Typeface: "Noto"}
				solars[29].Text = arr[29] + weeks[29] + "\n" + dgzarr[29] + "\n" + jcarr[29] + hharr[29] + "\n" + riQinarr[29]
				solars[29].TextSize = size
				return solars[29].Layout(gtx)
			}),
			space,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				solars[30].Font = text.Font{Typeface: "Noto"}
				solars[30].Text = arr[30] + weeks[30] + "\n" + dgzarr[30] + "\n" + jcarr[30] + hharr[30] + "\n" + riQinarr[30]
				solars[30].TextSize = size
				return solars[30].Layout(gtx)
			}),
			space,
		)
	})
}

var (
	day                                        int
	solars, jcs, ganzhis, riqins               []material.LabelStyle
	arr, weeks, dgzarr, jcarr, hharr, riQinarr []string
)

//
func (ui *UI) stars(t time.Time) {
	arr = nil
	weeks = nil
	dgzarr = nil
	jcarr = nil
	hharr = nil
	riQinarr = nil

	day = inits(t.Day())
	for i := 0; i < day; i++ {
		solar := material.Subtitle2(ui.theme, "")
		jcc := material.Subtitle2(ui.theme, "")
		moc := material.Subtitle2(ui.theme, "")
		rqc := material.Subtitle2(ui.theme, "")
		solars = append(solars, solar)
		jcs = append(jcs, jcc)
		ganzhis = append(ganzhis, moc)
		riqins = append(riqins, rqc)
		arr = append(arr, fmt.Sprintf("%d", i+1))
	}

	mgz := gz.MonthGZ(t.Year(), int(t.Month()), t.Day(), t.Hour())

	for i := 1; i <= day; i++ {
		dgz, _ := gz.DayGZ(t.Year(), int(t.Month()), i)
		dgzarr = append(dgzarr, dgz)
		jc := gz.JianChu(mgz, dgz)
		jcarr = append(jcarr, jc)
		hh := gz.GetRiHuangHei(mgz, dgz)
		hharr = append(hharr, hh)
		t = time.Date(y, time.Month(m), i, t.Hour(), 0, 0, 0, time.Local)

		riQin := gz.GetRiQin(int(t.Weekday()), dgz)
		riQinarr = append(riQinarr, riQin)
		week := pub.WeekName(int(t.Weekday()))
		weeks = append(weeks, week)
	}
}

func inits(days int) int {
	var b bool
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		b = true
	} else {
		b = false
	}
	switch int(t.Month()) {
	case 2:
		if b == true {
			days = 29
		} else {
			days = 28
		}
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	}
	return days
}

//
func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}

// func chineseFont() *material.Theme {
// 	f, err := os.Open("C:\\v2\\example\\font\\NotoSansSC-Regular.ttf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ttc, err := opentype.ParseCollectionReaderAt(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	th := material.NewTheme([]text.FontFace{{Face: ttc}})
// 	return th
// }
// func androidFont() *material.Theme {
// 	f, err := os.Open("/mnt/sdcard/Font/NotoSansSC-Regular.ttf")
// 	if err != nil {
// 		f, err := os.Create("/mnt/sdcard/Font/err.log")
// 		if err != nil {
// 			return nil
// 		}
// 		_, err = f.WriteString("no fonts")
// 		if err != nil {
// 			return nil
// 		}
// 		err = f.Close()
// 	}
// 	ttc, err := opentype.ParseCollectionReaderAt(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	th := material.NewTheme([]text.FontFace{{Face: ttc}})
// 	return th
// }
