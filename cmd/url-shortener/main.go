package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"

	"github.com/sanya-spb/URL-shortener/api/handler"
	"github.com/sanya-spb/URL-shortener/api/server"
	"github.com/sanya-spb/URL-shortener/app/repos/links"
	"github.com/sanya-spb/URL-shortener/app/starter"
	"github.com/sanya-spb/URL-shortener/db/memory/links/store"
)

var (
	lErr *log.Logger
	lOut *log.Logger
)

func main() {
	store := store.NewLinks()
	app, err := starter.NewApp(store)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if _, err := os.Stat(filepath.Dir(app.Config.LogAccess)); os.IsNotExist(err) {
		log.Fatalln(err.Error())
	}
	if _, err := os.Stat(filepath.Dir(app.Config.LogErrors)); os.IsNotExist(err) {
		log.Fatalln(err.Error())
	}
	if fAccess, err := os.OpenFile(app.Config.LogAccess, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalln(err.Error())
	} else {
		defer fAccess.Close()
		lOut = log.New(fAccess, "", log.LstdFlags)
		lOut.Println("run")
	}
	if fErrors, err := os.OpenFile(app.Config.LogErrors, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatalln(err.Error())
	} else {
		defer fErrors.Close()
		lErr = log.New(fErrors, "", log.LstdFlags)
		lErr.Println("run")
	}
	app.Welcome()

	links := links.NewLinks(store)
	hRouter := handler.NewRouter(links)
	srv := server.NewServer(app.Config.Listen, hRouter)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	log.Printf("listen at: %s\n", srv.Addr())
	go app.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
