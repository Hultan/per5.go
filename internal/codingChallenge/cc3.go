package codingChallenge

import (
	"math"

	"github.com/gotk3/gotk3/gdk"

	"github.com/hultan/per5/internal/per5"
)

type CC3 struct {
}

func newCC3() *CC3 {
	return &CC3{}
}

var s *snake
var scl = 20.0
var food per5.Vector2D

func (c *CC3) Setup(p *per5.Per5) {
	p.CreateCanvas(600, 600)
	p.FrameRate(10)
	p.SetKeyPressedFunc(c.KeyPressed)
	s = newSnake()
	food = p.CreateVector2D(p.Random(0, p.Width()), p.Random(0, p.Height()))
}

func (c *CC3) KeyPressed(p *per5.Per5) {
	switch p.KeyCode() {
	case gdk.KEY_Up:
		s.dir(0, -1)
	case gdk.KEY_Down:
		s.dir(0, 1)
	case gdk.KEY_Left:
		s.dir(-1, 0)
	case gdk.KEY_Right:
		s.dir(1, 0)
	}
}

func (c *CC3) Draw(p *per5.Per5) {
	// TODO : Continue here
	// p.Background(51)
	// s.update(p)
	// s.draw(p)
	//
	// s.eat(food)
	// p.FillRGBA(per5.ORANGE)
	// p.Rect(food.X, food.Y, scl, scl)
}

type snake struct {
	x, y   float64
	dx, dy float64
}

func newSnake() *snake {
	return &snake{0, 0, 1, 0}
}

func (s *snake) update(p *per5.Per5) {
	s.x += s.dx * scl
	s.y += s.dy * scl

	s.x = p.Constrain(s.x, 0, p.Width()-scl)
	s.y = p.Constrain(s.y, 0, p.Height()-scl)
}

func (s *snake) draw(p *per5.Per5) {
	p.Fill(255)
	p.Rect(s.x, s.y, scl, scl)
}

func (s *snake) dir(dx, dy float64) {
	s.dx = dx
	s.dy = dy
}

func (s *snake) eat(dx, dy float64) {
	s.dx = dx
	s.dy = dy
}

func pickLocation(p *per5.Per5) {
	cols := math.Floor(p.Width() / scl)
	rows := math.Floor(p.Height() / scl)
	food = p.CreateVector2D(p.Random(0, cols), p.Random(0, rows))
	food = food.Mult(scl)
}
