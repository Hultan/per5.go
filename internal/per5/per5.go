package per5

import (
	"image/color"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Per5 struct {
	win                    *gtk.ApplicationWindow
	da                     *gtk.DrawingArea
	ctx                    *cairo.Context
	setupFunc, drawFunc    func(*Per5)
	width, height          float64
	ticker                 ticker
	mouseX, mouseY         int
	strokeMode, fillMode   bool
	strokeColor, fillColor color.Color
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewDrawer(win *gtk.ApplicationWindow, da *gtk.DrawingArea, setup func(*Per5), draw func(*Per5)) *Per5 {
	d := &Per5{
		win:         win,
		da:          da,
		setupFunc:   setup,
		drawFunc:    draw,
		strokeMode:  true,
		fillMode:    true,
		strokeColor: BLACK,
		fillColor:   WHITE,
	}
	return d
}

func (p *Per5) Init() {
	// Events (signals)
	p.da.Connect("draw", p.onDraw)
	p.win.AddEvents(int(gdk.POINTER_MOTION_MASK))
	p.win.Connect("motion-notify-event", p.onMotionNotifyEvent)

	// Fields
	p.width, p.height = float64(p.da.GetAllocatedWidth()), float64(p.da.GetAllocatedHeight())

	// Ticker
	p.ticker.ticker = time.NewTicker(20 * time.Millisecond)
	p.ticker.tickerQuit = make(chan struct{})

	// Startup per5.go
	p.setupFunc(p)
	go p.mainLoop()
}

func (p *Per5) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	if p.ctx == nil {
		p.ctx = ctx
	}

	p.drawFunc(p)
}

func (p *Per5) onMotionNotifyEvent(da *gtk.ApplicationWindow, e *gdk.Event) {
	me := gdk.EventMotionNewFromEvent(e)

	x, y := me.MotionVal()
	xx, yy, err := p.win.TranslateCoordinates(p.da, int(x), int(y))
	if err != nil {
		panic(err)
	}
	p.mouseX, p.mouseY = xx, yy
}

func (p *Per5) mainLoop() {
	for {
		select {
		case <-p.ticker.ticker.C:
			p.da.QueueDraw()
		case <-p.ticker.tickerQuit:
			p.ticker.ticker.Stop()
			return
		}
	}
}
