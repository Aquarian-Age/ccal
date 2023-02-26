package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&UI{})
	myWindow := myApp.NewWindow("自定义布局")
	settings.NewSettings().LoadAppearanceScreen(myWindow)

	//st := NewSelectTime()
	//stcon := container.New(layout.NewGridLayout(5), st.ent, st.selectm, st.selectd, st.selecth, st.selectmin)
	//stBtnCon := container.New(layout.NewGridLayout(2), st.btnOK, st.btnAbout)
	//cont := container.New(&SelectTime{}, stcon, stBtnCon)

	ui := NewUI()
	//con4 := container.New(&G4{}, ui.G4.text, ui.G4.text1, ui.G4.text2, ui.G4.text3, ui.G4.btn)
	//con9 := container.New(&G9{}, ui.G9.text, ui.G9.text1, ui.G9.text2, ui.G9.text3, ui.G9.btn)
	//con2 := container.New(&G2{}, ui.G2.text, ui.G2.text1, ui.G2.text2, ui.G2.text3, ui.G2.btn)
	//conA := container.New(layout.NewHBoxLayout(), con4, con9, con2)
	//
	//con3 := container.New(&G3{}, ui.G3.text, ui.G3.text1, ui.G3.text2, ui.G3.text3, ui.G3.btn)
	//con5 := container.New(&G5{}, ui.G5.text, ui.G5.text1, ui.G5.text2, ui.G5.text3, ui.G5.btn)
	//con7 := container.New(&G7{}, ui.G7.text, ui.G7.text1, ui.G7.text2, ui.G7.text3, ui.G7.btn)
	//conB := container.New(layout.NewHBoxLayout(), con3, con5, con7)
	//
	//con8 := container.New(&G8{}, ui.G8.text, ui.G8.text1, ui.G8.text2, ui.G8.text3, ui.G8.btn)
	//con1 := container.New(&G1{}, ui.G1.text, ui.G1.text1, ui.G1.text2, ui.G1.text3, ui.G1.btn)
	//con6 := container.New(&G6{}, ui.G6.text, ui.G6.text1, ui.G6.text2, ui.G6.text3, ui.G6.btn)
	//conC := container.New(layout.NewHBoxLayout(), con8, con1, con6)
	//
	//conl := container.New(&Info{}, ui.Info.text)
	//
	//con := container.New(&UI{}, cont, conA, conB, conC, conl)
	//
	//st := ui.SelectTime
	//
	//ui.btnOK.OnTapped = func() {
	//	fmt.Println("btn clicked")
	//}
	//st.btnAbout.OnTapped = func() {
	//	ent := ui.ent
	//	year, _ := strconv.Atoi(ent.Text)
	//	month := ui.selectm.SelectedIndex() + 1
	//	day := ui.selectd.SelectedIndex() + 1
	//	hour := ui.selecth.SelectedIndex()
	//	min := ui.selectmin.SelectedIndex()
	//	fmt.Println("btn about: ", year, month, day, hour, min)
	//}
	//ui.G4.btn.OnTapped = func() {
	//	subW()
	//}
	////
	//ui.Info.text.ParseMarkdown(markdowns)
	//
	ui.OnBtnClicked()

	con := ui.con
	//
	myWindow.SetContent(con)
	myWindow.Resize(fyne.Size{Width: 380, Height: 640})
	myWindow.ShowAndRun()
}

var markdowns = `- a 
- b
- c
- d
- e
`

func (t *UI) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w = childSize.Width
		h = childSize.Height
	}
	return fyne.NewSize(w, h)

}
func (t *UI) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(t.MinSize(objects).Width, t.MinSize(objects).Height)
	pos = fyne.NewPos(0, t.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

var (
	white = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	black = &color.RGBA{A: 50}
	snow4 = &color.RGBA{R: 139, G: 137, B: 137, A: 150}
)

func (t *UI) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return white
	case theme.ColorNameButton, theme.ColorNameDisabled:
		return black
	case theme.ColorNamePlaceHolder, theme.ColorNameScrollBar:
		return black
	case theme.ColorNameInputBackground:
		return snow4
	case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
		return snow4
	case theme.ColorNameShadow:
		return snow4
	case theme.ColorNameForeground:
		return &color.RGBA{R: 0, G: 0, B: 0, A: 255}
	default:
		return snow4
	}
}

func (t *UI) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *UI) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
func (ui *UI) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		//指定返回的字体
		return resourceNotoSansSCRegularTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceNotoSansSCRegularTtf
}

func subW() {
	a := fyne.CurrentApp().NewWindow("sub windows")
	lab := widget.NewLabel("show sub windows")
	a.SetContent(lab)
	a.Resize(fyne.Size{Width: 600, Height: 600})
	a.Show()
}
