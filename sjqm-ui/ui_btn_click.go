package main

import (
	"fmt"
	"strconv"
)

func (ui *UI) OnBtnClicked() {

	st := ui.SelectTime

	ui.btnOK.OnTapped = func() {
		fmt.Println("btn clicked")
	}
	st.btnAbout.OnTapped = func() {
		ent := ui.ent
		year, _ := strconv.Atoi(ent.Text)
		month := ui.selectm.SelectedIndex() + 1
		day := ui.selectd.SelectedIndex() + 1
		hour := ui.selecth.SelectedIndex()
		min := ui.selectmin.SelectedIndex()
		fmt.Println("btn about: ", year, month, day, hour, min)
	}
	ui.G4.btn.OnTapped = func() {
		subW()
	}
	//
	ui.Info.text.ParseMarkdown(markdowns)
}
