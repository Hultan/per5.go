package draw

import (
	"github.com/gotk3/gotk3/gtk"
)

type Draw struct {
	da *gtk.DrawingArea
}

func NewDraw(da *gtk.DrawingArea) *Draw {
	return &Draw{da}
}

func (d *Draw) Setup() {

}