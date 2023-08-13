package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"project/backend/app"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}
	res.Write(fileData)
}

func main() {
	// Create an instance of the app structure
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "osu!-chart",
		Width:            1024,
		Height:           768,
		MinWidth:         400,
		MinHeight:        400,
		Frameless:        false,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			Theme:                             windows.Theme(windows.Dark),
			DisableFramelessWindowDecorations: true,
		},

		OnStartup: func(ctx context.Context) {
			defer func() {
				if p := recover(); p != nil {
					runtime.LogFatalf(ctx, "Application panicked: %v", p)
					runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
						Type:    runtime.ErrorDialog,
						Title:   "A fatal error occured!",
						Message: "Please try to restart the application, or consult the logs for more details.",
					})
				}
			}()

			app.Startup(ctx)
		},
		Bind: []interface{}{
			app,
		},
	},
	)

	if err != nil {
		println("Error:", err.Error())
	}
}
