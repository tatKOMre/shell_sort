package main

import (
	"shell/gui"
	"shell/sort"

	"fyne.io/fyne/v2/widget"
)

func OnSubmit(arr []int, label *widget.Label) {
	sort.Shellsort(&arr, label)
}

func main() {
	var gui *gui.Gui
	gui = gui.New("Сортировка Шелла")
	gui.NewInput("Введите чиcла", "Произвести сортировку", OnSubmit)
	gui.Window.ShowAndRun()
}
