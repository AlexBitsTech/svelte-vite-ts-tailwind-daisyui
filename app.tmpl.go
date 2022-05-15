package main

import (
	"context"
	"fmt"
	
	_ "github.com/wailsapp/wails/v2/pkg/runtime" // Remove underscore if package is needed
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// domready is called when all frontend elements are loaded.
func (a *App) domready(ctx context.Context) {
	fmt.Println("dom ready")
}


