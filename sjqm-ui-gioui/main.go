/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 *
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	w := app.NewWindow(app.Title("sjqm"),
		//app.MaxSize(unit.Dp(360), unit.Dp(640)))
		app.Size(unit.Dp(360), unit.Dp(640)))
	go func() {
		if err := loop(w); err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
func loop(w *app.Window) error {
	var ops op.Ops

	face, _ := opentype.Parse(resourceNotoSansSCRegularTtf)
	collection := []text.FontFace{{Font: text.Font{Typeface: "NotoSansSC-Regular"}, Face: face}}
	th := material.NewTheme(collection)
	res := SJQM(t)

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			strokeRect(gtx)
			strokeTriangle(gtx)
			strokeTriangle2(gtx)
			strokeTriangle3(gtx)
			strokeTriangle4(gtx)
			//输入信息布局
			res.layoutEditAndBtn(th, gtx)

			if res.btn1.Clicked() {
				y := convEdit(res.ed1.Text())
				m := convEdit(res.ed2.Text())
				d := convEdit(res.ed3.Text())
				h := convEdit(res.ed4.Text())
				if res.ed5.Text() == "" {
					res.ed5.SetText("0")
				}
				min := convEdit(res.ed5.Text())
				tx := time.Date(y, time.Month(m), d, h, min, 0, 0, time.Local)
				fmt.Println(tx.String()[:19])
				res = SJQM(tx)
			}
			if res.btna.Clicked() {
				res.btnGxs = abouts
			}
			if res.bg1.Clicked() {
				res.btnGxs = res.G1s()
			}
			if res.bg8.Clicked() {
				res.btnGxs = res.G8s()
			}
			if res.bg3.Clicked() {
				res.btnGxs = res.G3s()
			}
			if res.bg4.Clicked() {
				res.btnGxs = res.G4s()
			}
			if res.bg9.Clicked() {
				res.btnGxs = res.G9s()
			}
			if res.bg2.Clicked() {
				res.btnGxs = res.G2s()
			}
			if res.bg7.Clicked() {
				res.btnGxs = res.G7s()
			}
			if res.bg6.Clicked() {
				res.btnGxs = res.G6s()
			}

			res.layoutGanZhiS(gtx, th)

			res.show4(th, gtx)
			res.show9(th, gtx)
			res.show2(th, gtx)

			res.show3(th, gtx)
			res.show5(th, gtx)
			res.show7(th, gtx)

			res.show8(th, gtx)
			res.show1(th, gtx)
			res.show6(th, gtx)

			res.layoutGnInfo(th, gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}
}

var abouts = `时家奇门
V3.3.6
参考茅山起局方法。并进行了一定的修改。
mail:  
本软件仅限个人学习交流使用。
`
