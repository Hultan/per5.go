package draw_ui

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/per5/internal/codingChallenge"
	"github.com/hultan/per5/internal/per5"
	"github.com/hultan/softteam/framework"
)

const applicationTitle = "per5.go"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

var cc *codingChallenge.ChallengeManager

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
	title := fmt.Sprintf("%s %s - %s", applicationTitle, applicationVersion, applicationCopyRight)
	m.Window.SetTitle(title)

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Show the main window
	m.Window.ShowAll()

	da := m.builder.GetObject("drawingArea").(*gtk.DrawingArea)
	da.SetSizeRequest(600, 600)

	cc = codingChallenge.NewChallengeManager()
	cc.SetCurrentChallenge(0)

	d := per5.NewDrawer(m.Window, da, setup, draw)
	d.Init()
}

func setup(p *per5.Per5) {
	cc.Setup(p)
}

func draw(p *per5.Per5) {
	cc.Draw(p)
}
