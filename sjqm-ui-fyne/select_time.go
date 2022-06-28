package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

type SelectTime struct {
	ent                                  *widget.Entry
	selectm, selectd, selecth, selectmin *widget.Select
	btnOK, btnAbout                      *widget.Button
}

var t = time.Now().Local()

func NewSelectTime() *SelectTime {
	ent := widget.NewEntry()
	ent.SetText(strconv.Itoa(t.Year()))

	months := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	days := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	hours := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
	selectm := widget.NewSelect(months, func(s string) {
	})
	selectm.PlaceHolder = "month"
	selectd := widget.NewSelect(days, func(s string) {
	})
	selectd.PlaceHolder = "day"
	selecth := widget.NewSelect(hours, func(s string) {
	})
	selecth.PlaceHolder = "hour"

	mins := func() []string {
		var arr []string
		for i := 0; i < 60; i++ {
			s := fmt.Sprintf("%d", i)
			arr = append(arr, s)
		}
		return arr
	}
	selectmin := widget.NewSelect(mins(), func(s string) {})
	selectmin.PlaceHolder = "min"

	if selectm.SelectedIndex() == -1 {
		selectm.SetSelectedIndex(int(t.Month()) - 1)
	}
	if selectd.SelectedIndex() == -1 {
		selectd.SetSelectedIndex(t.Day() - 1)
	}
	if selecth.SelectedIndex() == -1 {
		selecth.SetSelectedIndex(t.Hour())
	}
	if selectmin.SelectedIndex() == -1 {
		selectmin.SetSelectedIndex(t.Minute())
	}

	btnok := widget.NewButton("btnok", func() {})
	btna := widget.NewButton("about", func() {})

	return &SelectTime{
		ent:       ent,
		selectm:   selectm,
		selectd:   selectd,
		selecth:   selecth,
		selectmin: selectmin,
		btnOK:     btnok,
		btnAbout:  btna,
	}
}
func (g SelectTime) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, g.MinSize(objects).Height-70)
	for _, o := range objects {
		size = o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(pos.X, size.Height))
	}
}

func (g SelectTime) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()
		w = childSize.Width
		//h = childSize.Height
	}
	return fyne.NewSize(w, h)
}
