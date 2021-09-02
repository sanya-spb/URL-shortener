package server

import (
	"context"
	"net/http"
	"time"

	"github.com/sanya-spb/URL-shortener/app/repos/links"
)

type Server struct {
	httpServer http.Server
	links      *links.Links
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.httpServer = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (srv *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	srv.httpServer.Shutdown(ctx)
	cancel()
}

func (srv *Server) Start(links *links.Links) {
	srv.links = links
	go srv.httpServer.ListenAndServe()
}

func (srv *Server) Addr() string {
	return srv.httpServer.Addr
}
