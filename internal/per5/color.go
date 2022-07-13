package per5

import (
	"image/color"
)

func (p *Per5) BackgroundRGBA(col color.Color) {
	p.setColorRGBA(col)
	p.drawBackground()
}

func (p *Per5) Background(col uint8) {
	p.setColor(col)
	p.drawBackground()
}

func (p *Per5) Fill(col uint8) {
	p.fillColor = color.RGBA{R: col, G: col, B: col, A: 255}
	p.fillMode = true
}

func (p *Per5) FillRGBA(col color.Color) {
	p.fillColor = col
	p.fillMode = true
}

func (p *Per5) NoFill() {
	p.fillMode = false
}

func (p *Per5) Stroke(col uint8) {
	p.strokeColor = color.RGBA{R: col, G: col, B: col, A: 255}
	p.strokeMode = true
}

func (p *Per5) StrokeRGBA(col color.Color) {
	p.strokeColor = col
	p.strokeMode = true
}

func (p *Per5) NoStroke() {
	p.strokeMode = false
}
