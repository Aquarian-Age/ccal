// SPDX-License-Identifier: Unlicense OR MIT

package main

// This projects demonstrates one way to manage and use multiple windows.
//
// It shows:
//   * how to track multiple windows,
//   * how to communicate between windows,
//   * how to create new windows.

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		a := NewApplication(ctx)
		letters := NewLetters()
		a.NewWindow("Letters", letters)
		a.Wait()
		os.Exit(0)
	}()

	app.Main()
}

// Application keeps track of all the windows and global state.
type Application struct {
	// Context is used to broadcast application shutdown.
	Context context.Context
	// Shutdown shuts down all windows.
	Shutdown func()
	// Theme is the application wide theme.
	Theme *material.Theme
	// active keeps track the open windows, such that application
	// can shut down, when all of them are closed.
	active sync.WaitGroup
}

func NewApplication(ctx context.Context) *Application {
	ctx, cancel := context.WithCancel(ctx)
	return &Application{
		Context:  ctx,
		Shutdown: cancel,

		Theme: material.NewTheme(gofont.Collection()),
	}
}

// Wait waits for all windows to close.
func (a *Application) Wait() {
	a.active.Wait()
}

// NewWindow creates a new tracked window.
func (a *Application) NewWindow(title string, view View, opts ...app.Option) {
	opts = append(opts, app.Title(title))
	w := &Window{
		App:    a,
		Window: app.NewWindow(opts...),
	}
	a.active.Add(1)
	go func() {
		defer a.active.Done()
		view.Run(w)
	}()
}

// Window holds window state.
type Window struct {
	App *Application
	*app.Window
}

// View describes .
type View interface {
	// Run handles the window event loop.
	Run(w *Window) error
}

// WidgetView allows to use layout.Widget as a view.
type WidgetView func(gtx layout.Context) layout.Dimensions

// Run displays the widget with default handling.
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
