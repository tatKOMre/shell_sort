package gui

import (
	"shell/sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SortType func(*[]int, *widget.Label)

func (g *Gui) NewInput(placeholdertitle string, buttontitle string, onSubmit func([]int, *widget.Label)) {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholdertitle)
	label := widget.NewLabel("")
	button := widget.NewButton(buttontitle, func() {
		arr := sort.Intar(input.Text)
		onSubmit(arr, label)
	})
	inputContainer := container.NewVBox(input, button)
	scrollContainer := container.NewScroll(label)
	scrollContainer.SetMinSize(fyne.NewSize(600, 300))
	content := container.NewVBox(inputContainer, scrollContainer)
	g.Window.SetContent(content)
	g.Window.Resize(fyne.NewSize(600, 400))
}
