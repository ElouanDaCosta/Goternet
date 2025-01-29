package drawing

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func Line() *canvas.Line {
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	return line
}
