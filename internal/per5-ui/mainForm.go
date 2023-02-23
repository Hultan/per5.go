package draw_ui

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/per5/internal/codingChallenge"
	"github.com/hultan/per5/internal/per5"
)

const applicationTitle = "per5.go"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

var cc *codingChallenge.ChallengeManager

//go:embed assets/main.glade
var mainGlade string

type MainForm struct {
	Window  *gtk.ApplicationWindow
	builder *gtkBuilder
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

	builder, err := newBuilder(mainGlade)
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.Window = m.builder.getObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	title := fmt.Sprintf("%s %s - %s", applicationTitle, applicationVersion, applicationCopyRight)
	m.Window.SetTitle(title)

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Show the main window
	m.Window.ShowAll()

	da := m.builder.getObject("drawingArea").(*gtk.DrawingArea)
	da.SetSizeRequest(600, 600)

	cc = codingChallenge.NewChallengeManager()
	cc.SetCurrentChallenge(0)

	d := per5.NewPer5(m.Window, da, cc.Setup, cc.Draw)
	d.Init()
}
