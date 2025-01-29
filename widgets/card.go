package widgets

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func Card() *widget.Card {
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{Y: 0x99}
	circle.StrokeWidth = 5
	newCard := widget.NewCard("Test", "", circle)
	return newCard
}
