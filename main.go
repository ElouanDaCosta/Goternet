package main

import (
	"goternet/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Goternet")

	// circle := drawing.Circle()
	// line := drawing.Line()
	card := widgets.Card()

	w.SetContent(card)

	w.Resize(fyne.NewSize(100, 100))

	w.ShowAndRun()
}
