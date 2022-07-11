package per5

import (
	"fmt"
	"image/color"
	"math"
)

func (d *Per5) CreateCanvas(width, height int) {
	d.da.SetSizeRequest(width, height)
	d.width, d.height = float64(width), float64(height)

	fmt.Println(d.da.GetAllocatedWidth(), "x", d.da.GetAllocatedHeight())
	fmt.Println(d.Width(), "x", d.Height())
}

func (d *Per5) BackgroundRGBA(col color.Color) {
	d.setColorRGBA(col)
	d.drawBackground()
}

func (d *Per5) Background(col uint8) {
	d.setColor(col)
	d.drawBackground()
}

func (d *Per5) Translate(dx, dy float64) {
	d.ctx.Translate(dx, dy)
	// d.translateX = dx
	// d.translateY = dy
}

func (d *Per5) Fill(col uint8) {
	d.fillColor = color.RGBA{R: col, G: col, B: col, A: 255}
	d.fillMode = true
}

func (d *Per5) FillRGBA(col color.Color) {
	d.fillColor = col
	d.fillMode = true
}

func (d *Per5) NoFill() {
	d.fillMode = false
}

func (d *Per5) Stroke(col uint8) {
	d.strokeColor = color.RGBA{R: col, G: col, B: col, A: 255}
	d.strokeMode = true
}

func (d *Per5) StrokeRGBA(col color.Color) {
	d.strokeColor = col
	d.strokeMode = true
}

func (d *Per5) NoStroke() {
	d.strokeMode = false
}

func (d *Per5) StrokeWeight(w float64) {
	d.ctx.SetLineWidth(w)
}

func (d *Per5) Width() float64 {
	return d.width
}

func (d *Per5) Height() float64 {
	return d.height
}

func (d *Per5) MouseX() int {
	return d.mouseX
}

func (d *Per5) MouseY() int {
	return d.mouseY
}

//
// Shapes
//

func (d *Per5) Point(x, y float64) {
	d.ctx.SetSourceRGBA(d.colorToGTK(d.strokeColor))
	d.Rect(d.xc(x), d.yc(y), 1, 1)
	d.ctx.Stroke()
}

func (d *Per5) Line(x1, y1, x2, y2 float64) {
	d.ctx.SetSourceRGBA(d.colorToGTK(d.strokeColor))
	d.ctx.MoveTo(d.xc(x1), d.yc(y1))
	d.ctx.LineTo(d.xc(x2), d.yc(y2))
	d.ctx.Stroke()
}

func (d *Per5) Square(x, y, s float64) {
	d.Rect(d.xc(x), d.yc(y), s, s)
}

func (d *Per5) Rect(x, y, w, h float64) {
	if d.fillMode {
		d.ctx.SetSourceRGBA(d.colorToGTK(d.fillColor))
		d.ctx.Rectangle(d.xc(x), d.yc(y), w, h)
		d.ctx.Fill()
	}
	if d.strokeMode {
		d.ctx.SetSourceRGBA(d.colorToGTK(d.strokeColor))
		d.ctx.Rectangle(d.xc(x), d.yc(y), w, h)
		d.ctx.Stroke()
	}
}

func (d *Per5) Circle(x, y, diam float64) {
	if d.fillMode {
		d.ctx.SetSourceRGBA(d.colorToGTK(d.fillColor))
		d.ctx.Arc(d.xc(x), d.yc(y), diam/2.0, 0, math.Pi*2)
		d.ctx.Fill()
	}
	if d.strokeMode {
		d.ctx.SetSourceRGBA(d.colorToGTK(d.strokeColor))
		d.ctx.Arc(d.xc(x), d.yc(y), diam/2.0, 0, math.Pi*2)
		d.ctx.Stroke()
	}
}

//
// Helper functions
//

func (d *Per5) drawBackground() {
	d.ctx.Rectangle(0, 0, d.width, d.height)
	d.ctx.Fill()
}

func (d *Per5) setColor(c uint8) {
	d.ctx.SetSourceRGBA(d.colorToGTK(color.RGBA{R: c, G: c, B: c, A: 255}))
}

func (d *Per5) setColorRGBA(c color.Color) {
	d.ctx.SetSourceRGBA(d.colorToGTK(c))
}

func (d *Per5) xc(x float64) float64 {
	return x + d.translateX
}

func (d *Per5) yc(y float64) float64 {
	return y + d.translateY
}

func (d *Per5) colorToGTK(c color.Color) (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}
