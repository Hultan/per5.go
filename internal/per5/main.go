package per5

import (
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Per5 struct {
	win                    *gtk.ApplicationWindow
	da                     *gtk.DrawingArea
	ctx                    *cairo.Context
	setupFunc              func(*Per5)
	drawFunc               func(*Per5)
	width, height          float64
	mode                   drawMode
	translateX, translateY float64
	ticker                 ticker
	mouseX, mouseY         int
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewDrawer(win *gtk.ApplicationWindow, da *gtk.DrawingArea, setup func(*Per5), draw func(*Per5)) *Per5 {
	d := &Per5{
		win:       win,
		da:        da,
		setupFunc: setup,
		drawFunc:  draw,
	}
	return d
}

func (d *Per5) Init() {
	// Events (signals)
	d.da.Connect("draw", d.onDraw)
	d.win.AddEvents(int(gdk.POINTER_MOTION_MASK))
	d.win.Connect("motion-notify-event", d.onMotionNotifyEvent)

	// Fields
	d.width, d.height = float64(d.da.GetAllocatedWidth()), float64(d.da.GetAllocatedHeight())

	// Ticker
	d.ticker.ticker = time.NewTicker(20 * time.Millisecond)
	d.ticker.tickerQuit = make(chan struct{})

	// Startup per5.go
	d.setupFunc(d)
	go d.mainLoop()
}

func (d *Per5) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	if d.ctx == nil {
		d.ctx = ctx
	}

	d.drawFunc(d)
}

func (d *Per5) onMotionNotifyEvent(da *gtk.ApplicationWindow, e *gdk.Event) {
	me := gdk.EventMotionNewFromEvent(e)

	x, y := me.MotionVal()
	xx, yy, err := d.win.TranslateCoordinates(d.da, int(x), int(y))
	if err != nil {
		panic(err)
	}
	d.mouseX, d.mouseY = xx, yy
}

func (d *Per5) mainLoop() {
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
