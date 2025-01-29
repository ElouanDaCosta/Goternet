package drawing

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func SetContentToText(str string) *canvas.Text {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText(str, green)
	text.TextStyle.Bold = true
	return text
}
