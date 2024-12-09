package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Gui struct {
	App    fyne.App
	Window fyne.Window
}

func (g *Gui) New(title string) *Gui {
	app := app.New()
	Window := app.NewWindow(title)
	return &Gui{
		App:    app,
		Window: Window,
	}
}
