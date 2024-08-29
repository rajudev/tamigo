package main

import {
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
}


func main() {
	a := app.New()
	w := a.NewWindow("germ")

	ui:= widget.NewTextGrid() // Creating a new Textgrid
	ui.SetText("This is a tamigo terminal!") // Set text to display


	// Creating a new Container with width 420 height 200

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			ui,
		),
	)

	w.ShowAndRun()
}