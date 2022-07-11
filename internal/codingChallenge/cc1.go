package codingChallenge

import (
	"math/rand"

	"github.com/hultan/per5/internal/per5"
)

type CC1 struct {
}

func newCC1() *CC1 {
	return &CC1{}
}

func (c *CC1) Setup(p *per5.Per5) {
	p.CreateCanvas(400, 400)
}

func (c *CC1) Draw(p *per5.Per5) {
	p.BackgroundRGBA(per5.GREY)
}

type Star struct {
	X, Y, Z float64
}

func newStar(width, height float64) *Star {
	return &Star{
		X: rand.Float64() * width,
		Y: rand.Float64() * height,
		Z: rand.Float64() * width,
	}
}

func (s *Star) Update() {

}
