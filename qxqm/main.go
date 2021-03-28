package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"gioui.org/font/opentype"
	"gioui.org/text"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		a := NewApplication(ctx)
		letters := NewLetters()
		a.NewWindow("日禽", letters)
		a.Wait()
		os.Exit(0)
	}()

	app.Main()
}

type Application struct {
	Context  context.Context
	Shutdown func()
	Theme    *material.Theme
	active   sync.WaitGroup
}

func NewApplication(ctx context.Context) *Application {
	ctx, cancel := context.WithCancel(ctx)
	return &Application{
		Context:  ctx,
		Shutdown: cancel,
		Theme:    utf8Font(),
	}
}

//utf8
func utf8Font() *material.Theme {
	f, err := os.Open("font/MaShanZheng_Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttf, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Println(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttf}})
	return th
}

func (a *Application) Wait() {
	a.active.Wait()
}

func (a *Application) NewWindow(title string, view View, opts ...app.Option) {
	opts = append(opts, app.Title(title))
	w := &Window{
		App:    a,
		Window: app.NewWindow(opts...),
	}
	a.active.Add(1)
	go func() {
		defer a.active.Done()
		_ = view.Run(w)
	}()
}

type Window struct {
	App *Application
	*app.Window
}

type View interface {
	Run(w *Window) error
}

type WidgetView func(gtx layout.Context) layout.Dimensions

func (view WidgetView) Run(w *Window) error {
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
