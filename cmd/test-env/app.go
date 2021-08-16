package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanya-spb/URL-shortener/internal/config"
	"github.com/sanya-spb/URL-shortener/pkg/version"
)

// application struct
type App struct {
	Version version.AppVersion
	Config  config.Config
	exPath  string
}

// init for App
func newApp() (*App, error) {
	var app *App = new(App)
	app.Version = *version.Version
	app.Config = *config.NewConfig()

	if ex, err := os.Executable(); err != nil {
		return nil, err
	} else {
		ex = filepath.Dir(ex)
		app.exPath = ex
	}
	return app, nil
}

// print welcome message
func (app *App) welcome() {
	fmt.Fprintf(os.Stdout, "Welcome to test!\nWorking directory: %s\nVersion: %s [%s@%s]\nCopyright: %s\n\n", app.exPath, app.Version.Version, app.Version.Commit, app.Version.BuildTime, app.Version.Copyright)
	fmt.Fprintf(os.Stdout, "Config dump: %+v\n", app.Config)
}
