package per5

import (
	"image/color"
	"math"
)

func (p *Per5) CreateCanvas(width, height int) {
	p.da.SetSizeRequest(width, height)
	p.width, p.height = float64(width), float64(height)
}

func (p *Per5) BackgroundRGBA(col color.Color) {
	p.setColorRGBA(col)
	p.drawBackground()
}

func (p *Per5) Background(col uint8) {
	p.setColor(col)
	p.drawBackground()
}

func (p *Per5) Translate(dx, dy float64) {
	p.ctx.Translate(dx, dy)
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

func (p *Per5) StrokeWeight(w float64) {
	p.ctx.SetLineWidth(w)
}

func (p *Per5) Width() float64 {
	return p.width
}

func (p *Per5) Height() float64 {
	return p.height
}

func (p *Per5) MouseX() int {
	return p.mouseX
}

func (p *Per5) MouseY() int {
	return p.mouseY
}

//
// Shapes
//

func (p *Per5) Point(x, y float64) {
	p.ctx.SetSourceRGBA(p.colorToGTK(p.strokeColor))
	p.Rect(x, y, 1, 1)
	p.ctx.Stroke()
}

func (p *Per5) Line(x1, y1, x2, y2 float64) {
	p.ctx.SetSourceRGBA(p.colorToGTK(p.strokeColor))
	p.ctx.MoveTo(x1, y1)
	p.ctx.LineTo(x2, y2)
	p.ctx.Stroke()
}

func (p *Per5) Square(x, y, s float64) {
	p.Rect(x, y, s, s)
}

func (p *Per5) RectMode(mode RectMode) {
	p.rectMode = mode
}

func (p *Per5) Rect(x, y, w, h float64) {
	switch p.rectMode {
	case RectMode_Corner:
	case RectMode_Corners:
		w -= x
		h -= y
	case RectMode_Center:
		x -= w / 2
		y -= h / 2
	case RectMode_Radius:
		x -= w
		y -= h
		w *= 2
		h *= 2
	}
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.fillColor))
		p.ctx.Rectangle(x, y, w, h)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.strokeColor))
		p.ctx.Rectangle(x, y, w, h)
		p.ctx.Stroke()
	}
}

func (p *Per5) Circle(x, y, diam float64) {
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.fillColor))
		p.ctx.Arc(x, y, diam/2.0, 0, math.Pi*2)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.strokeColor))
		p.ctx.Arc(x, y, diam/2.0, 0, math.Pi*2)
		p.ctx.Stroke()
	}
}

func (p *Per5) Triangle(x1, y1, x2, y2, x3, y3 float64) {
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.fillColor))
		p.Line(x1, y1, x2, y2)
		p.Line(x2, y2, x3, y3)
		p.Line(x3, y3, x1, y1)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToGTK(p.strokeColor))
		p.Line(x1, y1, x2, y2)
		p.Line(x2, y2, x3, y3)
		p.Line(x3, y3, x1, y1)
		p.ctx.Stroke()
	}
}

//
// Helper functions
//

func (p *Per5) drawBackground() {
	p.ctx.Rectangle(0, 0, p.width, p.height)
	p.ctx.Fill()
}

func (p *Per5) setColor(c uint8) {
	p.ctx.SetSourceRGBA(p.colorToGTK(color.RGBA{R: c, G: c, B: c, A: 255}))
}

func (p *Per5) setColorRGBA(c color.Color) {
	p.ctx.SetSourceRGBA(p.colorToGTK(c))
}

func (p *Per5) colorToGTK(c color.Color) (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}
