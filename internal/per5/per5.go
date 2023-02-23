package per5

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Per5 struct {
	core
	ticker                 ticker
	mouseX, mouseY         float64
	strokeMode, fillMode   bool
	strokeColor, fillColor color.Color
	rectMode               RectMode
	frameRate              float64
	frameCount             int
	keyCode                uint
}

type core struct {
	win                 *gtk.ApplicationWindow
	da                  *gtk.DrawingArea
	ctx                 *cairo.Context
	setupFunc, drawFunc func(*Per5)
	keyPressedFunc      func(*Per5)
	width, height       float64
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewPer5(win *gtk.ApplicationWindow, da *gtk.DrawingArea, setup func(*Per5), draw func(*Per5)) *Per5 {
	c := core{
		win:       win,
		da:        da,
		setupFunc: setup,
		drawFunc:  draw,
	}
	p := &Per5{core: c, frameRate: 60, frameCount: 0}
	p.resetPer5()
	rand.Seed(time.Now().UnixNano())

	return p
}

func (p *Per5) resetPer5() {
	p.strokeMode = true
	p.fillMode = true
	p.strokeColor = BLACK
	p.fillColor = WHITE
	p.rectMode = RectModeCorner
}

func (p *Per5) Init() {
	// Events (signals)
	p.da.Connect("draw", p.onDraw)
	p.win.AddEvents(int(gdk.POINTER_MOTION_MASK))
	p.win.Connect("motion-notify-event", p.onMotionNotifyEvent)
	p.win.Connect("key-press-event", p.onKeyDown)
	p.win.Connect("key-release-event", p.onKeyUp)

	// Fields
	p.width, p.height = float64(p.da.GetAllocatedWidth()), float64(p.da.GetAllocatedHeight())

	// Call setup
	p.setupFunc(p)

	// Ticker
	dur := time.Duration(1000.0 / p.frameRate)
	p.ticker.ticker = time.NewTicker(dur * time.Millisecond)
	p.ticker.tickerQuit = make(chan struct{})

	// Startup per5.go
	go p.mainLoop()
}

func (p *Per5) onDraw(_ *gtk.DrawingArea, ctx *cairo.Context) {
	// Save Cairo context if it is not saved yet
	if p.ctx == nil {
		p.ctx = ctx
	}

	// Reset and call the draw function again
	p.resetPer5()
	p.frameCount += 1
	p.drawFunc(p)
}

func (p *Per5) onMotionNotifyEvent(_ *gtk.ApplicationWindow, e *gdk.Event) {
	// Get mouseX and mouseY
	me := gdk.EventMotionNewFromEvent(e)
	x, y := me.MotionVal()
	xx, yy, err := p.win.TranslateCoordinates(p.da, int(x), int(y))
	if err != nil {
		panic(err)
	}
	p.mouseX, p.mouseY = float64(xx), float64(yy)
}

func (p *Per5) onKeyDown(_ *gtk.ApplicationWindow, e *gdk.Event) {
	if p.keyPressedFunc == nil {
		return
	}

	ke := gdk.EventKeyNewFromEvent(e)
	p.keyCode = ke.KeyVal()
	p.keyPressedFunc(p)
}

func (p *Per5) onKeyUp(_ *gtk.ApplicationWindow, e *gdk.Event) {
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
