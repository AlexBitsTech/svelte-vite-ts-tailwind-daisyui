package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "{{.ProjectName}}",
		Width:         900,
		Height:        600,
		DisableResize: false,
		Fullscreen:    false,
		Frameless:     false,
		// MinWidth:          400,
		// MinHeight:         400,
		// MaxWidth:          1280,
		// MaxHeight:         1024,
		StartHidden:       false,
		HideWindowOnClose: false,
		// RGBA:              &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		AlwaysOnTop:       false,
		Assets:            assets,
		// AssetsHandler:     assetsHandler,
		// Menu:              app.applicationMenu(),
		// Logger:            nil,
		// LogLevel:          logger.DEBUG,
		OnStartup:  app.startup,
		OnDomReady: app.domready,
		// OnShutdown:        app.shutdown,
		// OnBeforeClose:     app.beforeClose,
		WindowStartState: options.Maximised,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
