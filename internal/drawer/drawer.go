package drawer

import (
	"fmt"
	"image/color"
	"math"
)

func (d *Drawer) CreateCanvas(width, height int) {
	d.da.SetSizeRequest(width, height)
	d.width, d.height = float64(width), float64(height)

	fmt.Println(d.da.GetAllocatedWidth(), "x", d.da.GetAllocatedHeight())
	fmt.Println(d.Width(), "x", d.Height())
}

func (d *Drawer) BackgroundRGBA(col color.Color) {
	d.setColorRGBA(col)
	d.drawBackground()
}

func (d *Drawer) Background(col uint8) {
	d.setColor(col)
	d.drawBackground()
}

func (d *Drawer) Translate(dx, dy float64) {
	d.ctx.Translate(dx, dy)
	// d.translateX = dx
	// d.translateY = dy
}

func (d *Drawer) Fill(col uint8) {
	d.setColor(col)
	d.mode = drawModeFill
}

func (d *Drawer) FillRGBA(col color.Color) {
	d.setColorRGBA(col)
	d.mode = drawModeFill
}

func (d *Drawer) Stroke(col uint8) {
	d.setColor(col)
	d.mode = drawModeStroke
}

func (d *Drawer) StrokeRGBA(col color.Color) {
	d.setColorRGBA(col)
	d.mode = drawModeStroke
}

func (d *Drawer) StrokeWeight(w float64) {
	d.ctx.SetLineWidth(w)
}

func (d *Drawer) Width() float64 {
	return d.width
}

func (d *Drawer) Height() float64 {
	return d.height
}

//
// Shapes
//

func (d *Drawer) Point(x, y float64) {
	d.Rect(d.xc(x), d.yc(y), 1, 1)
	d.draw()
}

func (d *Drawer) Line(x1, y1, x2, y2 float64) {
	d.ctx.MoveTo(d.xc(x1), d.yc(y1))
	d.ctx.LineTo(d.xc(x2), d.yc(y2))
	d.draw()
}

func (d *Drawer) Square(x, y, s float64) {
	d.ctx.Rectangle(d.xc(x), d.yc(y), s, s)
	d.draw()
}

func (d *Drawer) Rect(x, y, w, h float64) {
	d.ctx.Rectangle(d.xc(x), d.yc(y), w, h)
	d.draw()
}

func (d *Drawer) Circle(x, y, diam float64) {
	d.ctx.Arc(d.xc(x), d.yc(y), diam/2.0, 0, math.Pi*2)
	d.draw()
}

//
// Helper functions
//

func (d *Drawer) drawBackground() {
	d.ctx.Rectangle(0, 0, d.width, d.height)
	d.ctx.Fill()
}

func (d *Drawer) draw() {
	switch d.mode {
	case drawModeFill:
		d.ctx.Fill()
	case drawModeStroke:
		d.ctx.Stroke()
	}
}

func (d *Drawer) setColor(c uint8) {
	d.ctx.SetSourceRGBA(d.colorToGTK(color.RGBA{R: c, G: c, B: c, A: 255}))
}

func (d *Drawer) setColorRGBA(c color.Color) {
	d.ctx.SetSourceRGBA(d.colorToGTK(c))
}

func (d *Drawer) xc(x float64) float64 {
	return x + d.translateX
}

func (d *Drawer) yc(y float64) float64 {
	return y + d.translateY
}

func (d *Drawer) colorToGTK(c color.Color) (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}
