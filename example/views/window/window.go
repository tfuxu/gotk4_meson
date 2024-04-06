package window

import (
	"github.com/tfuxu/gotk4_meson/example/constants"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type MainWindow struct {
	*gtk.ApplicationWindow
	settings *gio.Settings

	exampleButton *gtk.Button
}

func NewMainWindow(app *gtk.Application, settings *gio.Settings) *MainWindow {
	builder := gtk.NewBuilderFromResource(constants.RootPath + "/ui/main_window.ui")

	window := builder.GetObject("main_window").Cast().(*gtk.ApplicationWindow)
	window.SetApplication(app)
	window.SetDefaultSize(
		settings.Int("window-width"), settings.Int("window-height"),
	)

	button := builder.GetObject("example_button").Cast().(*gtk.Button)
	button.ConnectClicked(func() {
		println("Hi ^_^")
	})

	w := MainWindow{
		ApplicationWindow: window,
		settings:          settings,

		exampleButton:     button,
	}

	w.ConnectUnrealize(func() {
		saveWindowProps(window, settings)
	})

	return &w
}

func saveWindowProps(window *gtk.ApplicationWindow, settings *gio.Settings) {
	width, height := window.DefaultSize()

	settings.SetInt("window-width", width)
	settings.SetInt("window-height", height)
	settings.SetBoolean("window-maximized", window.IsMaximized())
}
