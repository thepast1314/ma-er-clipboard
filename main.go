////go:build windows && !production

package main

import (
	"cat-clipboard/backend"
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := backend.NewApp()
	applicationConfigContext := app.GetApplicationConfigContext()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "cat-clipboard",
		//Width:  1024,
		//Height: 768,
		Width:     applicationConfigContext.WindowConfig.Width,
		Height:    applicationConfigContext.WindowConfig.Height,
		MaxHeight: applicationConfigContext.WindowConfig.Height,
		MaxWidth:  applicationConfigContext.WindowConfig.Width,
		MinHeight: applicationConfigContext.WindowConfig.Height,
		MinWidth:  applicationConfigContext.WindowConfig.Width,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.Startup,
		Bind:             app.GetBeans(),
		// Windows 平台设置
		Windows: &windows.Options{
			// 窗口是否透明
			WebviewIsTransparent: false,
			// 窗口是否半透明
			WindowIsTranslucent: false,
			// 禁用窗口图标
			DisableWindowIcon: false,

			DisableFramelessWindowDecorations: true,
		},
		// Mac 平台设置
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "clipboard",
				Message: "A cross-platform clipboard manager",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
