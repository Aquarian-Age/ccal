package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type UI struct {
	SelectTime
	G4
	G9
	G2
	G3
	G5
	G7
	G8
	G1
	G6
	Info
	con *fyne.Container
}

func NewUI() *UI {
	st := NewSelectTime()
	g4 := &G4{
		btn:   widget.NewButton("btn4", func() {}),
		text:  canvas.NewText("八神 火", color.Black),
		text1: canvas.NewText("九星 土旺", color.Black),
		text2: canvas.NewText("八门 土生宫", color.Black),
		text3: canvas.NewText("暗干 4,5,3,8", color.Black),
	}
	g9 := &G9{
		btn:   widget.NewButton("btn9", func() {}),
		text:  canvas.NewText("八神 火9", color.Black),
		text1: canvas.NewText("九星 土旺9", color.Black),
		text2: canvas.NewText("八门 土生宫9", color.Black),
		text3: canvas.NewText("暗干 9,3,2,7", color.Black),
	}
	g2 := &G2{
		btn:   widget.NewButton("btn2", func() {}),
		text:  canvas.NewText("八神 火2", color.Black),
		text1: canvas.NewText("九星 土旺2", color.Black),
		text2: canvas.NewText("八门 土生宫2", color.Black),
		text3: canvas.NewText("暗干 2,8,5,10", color.Black),
	}
	g3 := &G3{
		btn:   widget.NewButton("btn3", func() {}),
		text:  canvas.NewText("八神 火3", color.Black),
		text1: canvas.NewText("九星 土旺3", color.Black),
		text2: canvas.NewText("八门 土生宫3", color.Black),
		text3: canvas.NewText("暗干 3,4,3,8", color.Black),
	}
	g5 := &G5{
		btn:   widget.NewButton("btn5", func() {}),
		text:  canvas.NewText("八神 火5", color.Black),
		text1: canvas.NewText("九星 土旺5", color.Black),
		text2: canvas.NewText("八门 土生宫5", color.Black),
		text3: canvas.NewText("暗干 5", color.Black),
	}
	g7 := &G7{
		btn:   widget.NewButton("btn7", func() {}),
		text:  canvas.NewText("八神 火7", color.Black),
		text1: canvas.NewText("九星 土旺7", color.Black),
		text2: canvas.NewText("八门 土生宫7", color.Black),
		text3: canvas.NewText("暗干 7,2,4,9 ", color.Black),
	}
	g8 := &G8{
		btn:   widget.NewButton("btn8", func() {}),
		text:  canvas.NewText("八神 火8", color.Black),
		text1: canvas.NewText("九星 土旺8", color.Black),
		text2: canvas.NewText("八门 土生宫8", color.Black),
		text3: canvas.NewText("暗干 8,8,5,10 ", color.Black),
	}
	g1 := &G1{
		btn:   widget.NewButton("btn1", func() {}),
		text:  canvas.NewText("八神 火1", color.Black),
		text1: canvas.NewText("九星 土旺1", color.Black),
		text2: canvas.NewText("八门 土生宫1", color.Black),
		text3: canvas.NewText("暗干 1,6,1,6 ", color.Black),
	}
	g6 := &G6{
		btn:   widget.NewButton("btn6", func() {}),
		text:  canvas.NewText("八神 火6", color.Black),
		text1: canvas.NewText("九星 土旺6", color.Black),
		text2: canvas.NewText("八门 土生宫6", color.Black),
		text3: canvas.NewText("暗干 6,1,4,10 ", color.Black),
	}

	info := &Info{
		widget.NewRichTextWithText(`- richtext
- a
- b
- aaa
- ccc`),
	}
	//
	stcon := container.New(layout.NewGridLayout(5), st.ent, st.selectm, st.selectd, st.selecth, st.selectmin)
	stBtnCon := container.New(layout.NewGridLayout(2), st.btnOK, st.btnAbout)
	cont := container.New(&SelectTime{}, stcon, stBtnCon)

	con4 := container.New(&G4{}, g4.text, g4.text1, g4.text2, g4.text3, g4.btn)
	con9 := container.New(&G9{}, g9.text, g9.text1, g9.text2, g9.text3, g9.btn)
	con2 := container.New(&G2{}, g2.text, g2.text1, g2.text2, g2.text3, g2.btn)
	conA := container.New(layout.NewHBoxLayout(), con4, con9, con2)

	con3 := container.New(&G3{}, g3.text, g3.text1, g3.text2, g3.text3, g3.btn)
	con5 := container.New(&G5{}, g5.text, g5.text1, g5.text2, g5.text3, g5.btn)
	con7 := container.New(&G7{}, g7.text, g7.text1, g7.text2, g7.text3, g7.btn)
	conB := container.New(layout.NewHBoxLayout(), con3, con5, con7)

	con8 := container.New(&G8{}, g8.text, g8.text1, g8.text2, g8.text3, g8.btn)
	con1 := container.New(&G1{}, g1.text, g1.text1, g1.text2, g1.text3, g1.btn)
	con6 := container.New(&G6{}, g6.text, g6.text1, g6.text2, g6.text3, g6.btn)
	conC := container.New(layout.NewHBoxLayout(), con8, con1, con6)

	conl := container.New(&Info{}, info.text)

	con := container.New(&UI{}, cont, conA, conB, conC, conl)

	return &UI{
		SelectTime: *st,
		G4:         *g4,
		G9:         *g9,
		G2:         *g2,
		G3:         *g3,
		G5:         *g5,
		G7:         *g7,
		G8:         *g8,
		G1:         *g1,
		G6:         *g6,
		Info:       *info,
		con:        con,
	}
}

type G4 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G4) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G4) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 30
	}
	return fyne.NewSize(w, h)
}

type G9 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G9) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G9) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 30
	}
	return fyne.NewSize(w, h)
}

type G2 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G2) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G2) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 30
	}
	return fyne.NewSize(w, h)
}

type G3 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G3) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G3) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 10
	}
	return fyne.NewSize(w, h)
}

type G5 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G5) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G5) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 10
	}
	return fyne.NewSize(w, h)
}

type G7 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G7) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G7) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height - 10
	}
	return fyne.NewSize(w, h)
}

type G8 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G8) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100*2)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G8) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height
	}
	return fyne.NewSize(w, h)
}

type G1 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G1) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100*2)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G1) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height
	}
	return fyne.NewSize(w, h)
}

type G6 struct {
	btn   *widget.Button
	text  *canvas.Text
	text1 *canvas.Text
	text2 *canvas.Text
	text3 *canvas.Text
}

func (g G6) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(g.MinSize(objects).Width, g.MinSize(objects).Height)
	pos = fyne.NewPos(0, g.MinSize(objects).Height+100*2)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g G6) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for k, o := range objects {
		childSize := o.MinSize()
		if k == 4 {
			w = childSize.Width + 60
		} else {
			w = childSize.Width
		}

		h = childSize.Height
	}
	return fyne.NewSize(w, h)
}

type Info struct {
	text *widget.RichText
}

func (g Info) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, g.MinSize(objects).Height+100*2+90)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g Info) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w = childSize.Width
		h = childSize.Height - 60
	}
	return fyne.NewSize(w, h)
}
