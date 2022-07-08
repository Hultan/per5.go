package drawer

import (
	"image/color"
)

type drawMode uint8

const (
	drawModeFill drawMode = iota
	drawModeStroke
)

var WHITE = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var BLACK = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var RED = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var GREEN = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var BLUE = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var ORANGE = color.RGBA{R: 255, G: 165, B: 0, A: 255}
var YELLOW = color.RGBA{R: 255, G: 255, B: 0, A: 255}
var INDIGO = color.RGBA{R: 75, G: 0, B: 130, A: 255}
var VIOLET = color.RGBA{R: 238, G: 130, B: 238, A: 255}
var BROWN = color.RGBA{R: 118, G: 57, B: 49, A: 255}
