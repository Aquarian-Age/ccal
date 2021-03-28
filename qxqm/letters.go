package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Letters struct {
	win   *Window
	items []*LetterListItem
	list  layout.List
}
type LetterListItem struct {
	Number int
	Text   string
	Click  widget.Clickable
}

func NewLetters() *Letters {
	view := &Letters{
		list: layout.List{Axis: layout.Vertical},
	}
	for i := 0; i < len(starNames); i++ {
		view.items = append(view.items, &LetterListItem{Text: starNames[i]})
	}
	return view
}

func (v *Letters) Run(w *Window) error {
	v.win = w
	return wly(v.LayoutN).Run(w)
}

var (
	starNames = []string{
		"角木蛟", "亢金龙", "氐土貉", "房日兔", "心月狐", "尾火虎", "箕水豹", //东方青龙
		"斗木獬", "牛金牛", "女土蝠", "虚日鼠", "危月燕", "室火猪", "壁水貐", //北方玄武
		"奎木狼", "娄金狗", "胃土雉", "昴日鸡", "畢月乌", "觜火猴", "参水猿", //西方白虎
		"井木犴", "鬼金羊", "柳土獐", "星日马", "张月鹿", "翼火蛇", "轸水蚓", //南方朱雀
	}
)

func bfyqRQ(name string) {
	go func() {
		w := app.NewWindow(app.Size(unit.Dp(1600), unit.Dp(850)))
		events := w.Events()
		var ops op.Ops
		for {
			e := <-events
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				ns := name + ".png"
				showPNG(ns, gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
}
