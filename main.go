package main

import (
	"embed"
	"todo-app/internal/db"
	"todo-app/internal/service"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	todoService := service.NewTodoService(db.NewInMemoryDB())

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Todos",
		Width:  1024,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			todoService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
