package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sanya-spb/URL-shortener/api/auth"
	"github.com/sanya-spb/URL-shortener/api/handler"
)

type Router struct {
	http.Handler
	hHandler *handler.Handler
}

func NewRouter(hHandler *handler.Handler) *Router {
	rRouter := &Router{
		hHandler: hHandler,
	}

	r := chi.NewRouter()
	r.Group(func(rAdm chi.Router) {
		rAdm.Get("/stat", rRouter.Stat)
		rAdm.Put("/update", rRouter.Update)
		rAdm.Delete("/delete", rRouter.Delete)
	}).Use(auth.AuthMiddleware) // TODO: заменить на middleware.BasicAuth()
	r.Put("/new", rRouter.Create)
	r.Get("/info", rRouter.Read)
	// r.Put("/update", rRouter.Update)
	// r.Delete("/delete", rRouter.Delete)
	r.Get("/go", rRouter.Go)
	// r.Get("/stat", rRouter.Stat)

	return rRouter
}

func (rRouter *Router) Create(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Read(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Update(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Delete(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Go(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Stat(w http.ResponseWriter, r *http.Request)
