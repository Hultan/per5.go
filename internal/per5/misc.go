package per5

import (
	"image/color"
)

//
// Helper functions
//

func (p *Per5) drawBackground() {
	p.ctx.Rectangle(0, 0, p.width, p.height)
	p.ctx.Fill()
}

func (p *Per5) setColor(c uint8) {
	p.ctx.SetSourceRGBA(p.colorToRGBA(color.RGBA{R: c, G: c, B: c, A: 255}))
}

func (p *Per5) setColorRGBA(c color.Color) {
	p.ctx.SetSourceRGBA(p.colorToRGBA(c))
}

func (p *Per5) colorToRGBA(c color.Color) (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}
