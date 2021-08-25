package main

import (
	"github.com/andlabs/ui"
)

type goUI struct {
	child  *ui.Box
	window *ui.Window
}

func (g goUI) Start() {
	ui.Main(func() {
		if g.window == nil {
			g.window = ui.NewWindow("Default goUI title", 640, 480, true)
		}

		g.window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		ui.OnShouldQuit(func() bool {
			g.window.Destroy()
			return true
		})

		if g.child != nil {
			g.window.SetChild(g.child)
		}

		g.window.Show()
	})
}

func (g goUI) Window(title string, width int, height int, hasMenubar bool) *goUI {
	g.window = ui.NewWindow(title, width, height, hasMenubar)
	return &g
}

func App(child *ui.Box) *goUI {
	ui := goUI{}
	ui.child = child

	return &ui
}
