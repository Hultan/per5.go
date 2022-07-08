package drawer

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Drawer struct {
	win                    *gtk.ApplicationWindow
	da                     *gtk.DrawingArea
	ctx                    *cairo.Context
	setupFunc              func(*Drawer)
	drawFunc               func(*Drawer)
	width, height          float64
	mode                   drawMode
	translateX, translateY float64
	ticker                 ticker
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewDrawer(win *gtk.ApplicationWindow, da *gtk.DrawingArea, setup func(*Drawer), draw func(*Drawer)) *Drawer {

	d := &Drawer{
		win:       win,
		da:        da,
		setupFunc: setup,
		drawFunc:  draw,
	}
	return d
}

func (d *Drawer) Init() {
	d.width, d.height = float64(d.da.GetAllocatedWidth()), float64(d.da.GetAllocatedHeight())
	d.da.Connect("draw", d.onDraw)
	d.win.AddEvents(int(gdk.POINTER_MOTION_MASK))
	d.win.Connect("motion-notify-event", d.onMotionNotifyEvent)

	d.ticker.ticker = time.NewTicker(20 * time.Millisecond)
	d.ticker.tickerQuit = make(chan struct{})

	d.setupFunc(d)

	go d.mainLoop()
}

func (d *Drawer) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	if d.ctx == nil {
		d.ctx = ctx
	}

	d.drawFunc(d)
}

func (d *Drawer) mainLoop() {
	for {
		select {
		case <-d.ticker.ticker.C:
			d.da.QueueDraw()
		case <-d.ticker.tickerQuit:
			d.ticker.ticker.Stop()
			return
		}
	}
}

func (d *Drawer) onMotionNotifyEvent(da *gtk.ApplicationWindow, e *gdk.Event) {
	me := gdk.EventMotionNewFromEvent(e)

	x, y := me.MotionVal()
	xx, yy, err := d.win.TranslateCoordinates(d.da, int(x), int(y))
	if err != nil {
		panic(err)
	}
	fmt.Println(xx, ",", yy)
}
