package codingChallenge

import (
	"fmt"

	"github.com/hultan/per5/internal/per5"
)

type CC1 struct {
}

func newCC1() *CC1 {
	return &CC1{}
}

func (c *CC1) Setup(p *per5.Per5) {
	p.CreateCanvas(800, 800)
	p.FrameRate(10)
	for i := 0; i < 800; i++ {
		stars = append(stars, newStar(p))
	}
}

func (c *CC1) Draw(p *per5.Per5) {
	speed = p.Map(p.MouseX(), 0, p.Width(), 0, 50)
	p.BackgroundRGBA(per5.BLACK)
	p.Translate(p.Width()/2, p.Height()/2)

	for _, s := range stars {
		s.update(p)
		s.draw(p)
	}

	fmt.Println("Frames :", p.FrameCount())
}

var stars []*star
var speed float64

type star struct {
	X, Y, Z float64
	PZ      float64
}

func newStar(p *per5.Per5) *star {
	s := &star{
		X: p.Random(-p.Width(), p.Width()),
		Y: p.Random(-p.Height(), p.Height()),
		Z: p.Random(1, p.Width()),
	}
	s.PZ = s.Z

	return s
}

func (s *star) update(p *per5.Per5) {
	s.Z -= speed
	if s.Z < 1 {
		s.X = p.Random(-p.Width(), p.Width())
		s.Y = p.Random(-p.Height(), p.Height())
		s.Z = p.Random(1, p.Width())
		s.PZ = s.Z
	}
}

func (s *star) draw(p *per5.Per5) {
	p.Fill(255)
	p.NoStroke()

	sx := p.Map(s.X/s.Z, 0, 1, 0, p.Width())
	sy := p.Map(s.Y/s.Z, 0, 1, 0, p.Height())

	// p.Circle(sx, sy, p.Map(s.Z, 0, p.Width(), 16, 1))

	px := p.Map(s.X/s.PZ, 0, 1, 0, p.Width())
	py := p.Map(s.Y/s.PZ, 0, 1, 0, p.Height())

	s.PZ = s.Z

	p.Stroke(255)
	p.Line(px, py, sx, sy)
}
