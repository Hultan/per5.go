package draw_ui

import (
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/draw/internal/per5"
	"github.com/hultan/softteam/framework"
)

const applicationTitle = "drawer"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window  *gtk.ApplicationWindow
	builder *framework.GtkBuilder
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.Window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("drawer-ui main window")

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Show the main window
	m.Window.ShowAll()

	da := m.builder.GetObject("drawingArea").(*gtk.DrawingArea)
	da.SetSizeRequest(600, 600)
	d := per5.NewDrawer(m.Window, da, setup, draw)
	d.Init()
}

var x = 0.0
var y = 0.0
var dim = 80.0

func setup(d *per5.Per5) {
	d.CreateCanvas(720, 400)
}

func draw(d *per5.Per5) {
	// d.Background(102)
	// // Animate by increasing our x value
	// x = x + 0.8
	// // If the shape goes off the canvas, reset the position
	// if x > d.Width()+dim {
	// 	x = -dim
	// }
	//
	// // Even though our rect command draws the shape with its
	// // center at the origin, translate moves it to the new
	// // x and y position
	// d.Translate(x, d.Height()/2-dim/2)
	// d.Fill(255)
	// d.Rect(-dim/2, -dim/2, dim, dim)
	//
	// // Transforms accumulate. Notice how this rect moves
	// // twice as fast as the other, but it has the same
	// // parameter for the x-axis value
	// d.Translate(x, dim)
	// d.Fill(0)
	// d.Rect(-dim/2, -dim/2, dim, dim)

	d.Translate(100, 100)
	d.Rect(50, 50, 200, 200)
}
