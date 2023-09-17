package main

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/tfuxu/gotk4_meson/example/constants"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	resources, error := gio.ResourceLoad(filepath.Join(constants.PkgDataDir, "gotk4_meson.gresource"))
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	gio.ResourcesRegister(resources)

	settings := gio.NewSettings(constants.AppID)

	app := gtk.NewApplication(constants.AppID, gio.ApplicationFlagsNone)
	app.SetResourceBasePath(constants.RootPath)
	app.ConnectActivate(func() {
		activate(app, settings)
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application, settings *gio.Settings) {
	builder := gtk.NewBuilderFromResource(constants.RootPath + "/ui/main_window.ui")

	window := builder.GetObject("main_window").Cast().(*gtk.ApplicationWindow)
	window.SetApplication(app)
	window.SetDefaultSize(
		settings.Int("window-width"), settings.Int("window-height"),
	)

	window.ConnectUnrealize(func() {
		saveWindowProps(window, settings)
	})

	button := builder.GetObject("example_button").Cast().(*gtk.Button)
	button.ConnectClicked(func() {
		println("Hi ^_^")
	})

	window.Show()
}

func saveWindowProps(window *gtk.ApplicationWindow, settings *gio.Settings) {
	width, height := window.DefaultSize()

	settings.SetInt("window-width", width)
	settings.SetInt("window-height", height)
	settings.SetBoolean("window-maximized", window.IsMaximized())
}
