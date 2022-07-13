package per5

import (
	"math"
)

//
// Shapes
//

func (p *Per5) Point(x, y float64) {
	p.ctx.SetSourceRGBA(p.colorToRGBA(p.strokeColor))
	p.Rect(x, y, 1, 1)
	p.ctx.Stroke()
}

func (p *Per5) Line(x1, y1, x2, y2 float64) {
	p.ctx.SetSourceRGBA(p.colorToRGBA(p.strokeColor))
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
	case RectModeCorner:
	case RectModeCorners:
		w -= x
		h -= y
	case RectModeCenter:
		x -= w / 2
		y -= h / 2
	case RectModeRadius:
		x -= w
		y -= h
		w *= 2
		h *= 2
	}
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.fillColor))
		p.ctx.Rectangle(x, y, w, h)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.strokeColor))
		p.ctx.Rectangle(x, y, w, h)
		p.ctx.Stroke()
	}
}

func (p *Per5) Circle(x, y, diam float64) {
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.fillColor))
		p.ctx.Arc(x, y, diam/2.0, 0, math.Pi*2)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.strokeColor))
		p.ctx.Arc(x, y, diam/2.0, 0, math.Pi*2)
		p.ctx.Stroke()
	}
}

func (p *Per5) Triangle(x1, y1, x2, y2, x3, y3 float64) {
	if p.fillMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.fillColor))
		p.Line(x1, y1, x2, y2)
		p.Line(x2, y2, x3, y3)
		p.Line(x3, y3, x1, y1)
		p.ctx.Fill()
	}
	if p.strokeMode {
		p.ctx.SetSourceRGBA(p.colorToRGBA(p.strokeColor))
		p.Line(x1, y1, x2, y2)
		p.Line(x2, y2, x3, y3)
		p.Line(x3, y3, x1, y1)
		p.ctx.Stroke()
	}
}

func (p *Per5) StrokeWeight(w float64) {
	p.ctx.SetLineWidth(w)
}
