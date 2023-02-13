package main

import (
	"embed"

	"github.com/ppz-pro/ppz/backend/api"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	runtime := &api.Runtime{}
	app := &api.App{}

	// Create application
	err := wails.Run(&options.App{
		Title:  "皮皮仔",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: runtime.Startup,
		Bind: []interface{}{
			runtime,
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
