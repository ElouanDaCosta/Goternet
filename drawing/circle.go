package drawing

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

func Circle() *canvas.Circle {
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Gray{Y: 0x99}
	circle.StrokeWidth = 5
	return circle
}
