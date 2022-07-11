package codingChallenge

import (
	"github.com/hultan/per5/internal/per5"
)

type CC0 struct {
}

func newCC0() *CC0 {
	return &CC0{}
}

func (c *CC0) Setup(p *per5.Per5) {
	p.CreateCanvas(600, 600)
}

func (c *CC0) Draw(p *per5.Per5) {
	p.BackgroundRGBA(per5.GREY)
	p.StrokeRGBA(per5.WHITE)
	p.NoFill()

	// Horizontal lines
	p.Line(0, 100, p.Width(), 100)
	p.Line(0, 200, p.Width(), 200)
	p.Line(0, 300, p.Width(), 300)

	// Vertical lines
	p.Line(100, 0, 100, p.Height())
	p.Line(200, 0, 200, p.Height())
	p.Line(300, 0, 300, p.Height())

	p.StrokeRGBA(per5.RED)
	p.Rect(100, 100, 100, 100)

	p.StrokeRGBA(per5.BLUE)
	p.RectMode(per5.RectMode_Corners)
	p.Rect(100, 100, 150, 150)

	p.StrokeRGBA(per5.BLACK)
	p.RectMode(per5.RectMode_Center)
	p.Rect(175, 175, 100, 100)

	p.StrokeRGBA(per5.GREEN)
	p.RectMode(per5.RectMode_Radius)
	p.Rect(175, 175, 100, 100)

	p.Triangle(50, 50, 150, 50, 300, 100)

	// p.Translate(p.Width()/2, p.Height()/2)
	// p.Stroke(255)
	//
	// p.Line(-100, -100, -100, 100)
	// p.Line(-100, 100, 100, 100)
	// p.Line(100, 100, 100, -100)
	// p.Line(100, -100, -100, -100)
	//
	// p.Line(-100/1.2, -100/1.2, -100/1.2, 100/1.2)
	// p.Line(-100/1.2, 100/1.2, 100/1.2, 100/1.2)
	// p.Line(100/1.2, 100/1.2, 100/1.2, -100/1.2)
	// p.Line(100/1.2, -100/1.2, -100/1.2, -100/1.2)
	//
	// p.Line(-100, -100, -100/1.2, -100/1.2)
	// p.Line(-100, 100, -100/1.2, 100/1.2)
	// p.Line(100, 100, 100/1.2, 100/1.2)
	// p.Line(100, -100, 100/1.2, -100/1.2)
}
