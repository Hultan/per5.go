package codingChallenge

import (
	"github.com/hultan/per5/internal/per5"
)

type CC1 struct {
}

func newCC1() *CC1 {
	return &CC1{}
}

func (c *CC1) Setup(p *per5.Per5) {
	p.CreateCanvas(600, 600)
}

func (c *CC1) Draw(p *per5.Per5) {
	p.BackgroundRGBA(per5.BLACK)
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

type Star struct {
	X, Y, Z float64
}

// func (c *CC1)
