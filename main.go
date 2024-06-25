package main

import (
	"embed"
	"github.com/run-bigpig/xpvideo/internal/bootstrap"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := bootstrap.NewBoot()
	err := wails.Run(&options.App{
		MinWidth:  1440,
		MinHeight: 768,
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app.NewController(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
