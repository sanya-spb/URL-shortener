package router

import (
	"net/http"

	"github.com/sanya-spb/URL-shortener/api/handler"
)

type Router struct {
	http.Handler
	hs *handler.Handler
}

func NewRouter(hs *handler.Handler) *Router {
	rRouter := &Router{
		hs: hs,
	}
	return rRouter
}
