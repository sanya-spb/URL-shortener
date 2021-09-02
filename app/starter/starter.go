package starter

import (
	"context"
	"log"
	"sync"

	"github.com/sanya-spb/URL-shortener/app/repos/links"
	"github.com/sanya-spb/URL-shortener/internal/config"
	"github.com/sanya-spb/URL-shortener/pkg/version"
)

// application struct
type App struct {
	Links   *links.Links
	Version version.AppVersion
	Config  config.Config
}

// init for App
func NewApp(store links.LinksStore) (*App, error) {
	app := &App{
		Links:   links.NewLinks(store),
		Version: *version.Version,
		Config:  *config.NewConfig(),
	}
	return app, nil
}

type HTTPServer interface {
	Start(links *links.Links)
	Stop()
}

// start service
func (app *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(app.Links)
	<-ctx.Done()
	hs.Stop()
}

// print welcome message
func (app *App) Welcome() {
	log.Printf("Starting url-shortener!\n\nVersion: %s [%s@%s]\nCopyright: %s\n\n", app.Version.Version, app.Version.Commit, app.Version.BuildTime, app.Version.Copyright)
	log.Printf("Config dump: %+v\n", app.Config)
}
