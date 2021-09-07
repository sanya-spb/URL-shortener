package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sanya-spb/URL-shortener/api/auth"
	"github.com/sanya-spb/URL-shortener/api/handler"
)

type Router struct {
	http.Handler
	hHandler *handler.Handler
}

type TLink handler.TLink

func (TLink) Bind(r *http.Request) error
func (TLink) Render(w http.ResponseWriter, r *http.Request) error

func NewRouter(hHandler *handler.Handler) *Router {
	rRouter := &Router{
		hHandler: hHandler,
	}

	r := chi.NewRouter()
	r.Group(func(rAdm chi.Router) {
		rAdm.Get("/stat", rRouter.Stat)
		rAdm.Put("/update", rRouter.Update)
		rAdm.Delete("/delete/{uuid}", rRouter.Delete)
	}).Use(auth.AuthMiddleware) // TODO: заменить на middleware.BasicAuth()
	r.Put("/new", rRouter.Create)
	r.Get("/info/{uuid}", rRouter.Read)
	// r.Put("/update", rRouter.Update)
	// r.Delete("/delete", rRouter.Delete)
	r.Get("/go/{id}", rRouter.Go)
	// r.Get("/stat", rRouter.Stat)

	return rRouter
}

func (rRouter *Router) Create(w http.ResponseWriter, r *http.Request) {
	rl := TLink{}
	if err := render.Bind(r, &rl); err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	l, err := rRouter.hHandler.Create(r.Context(), handler.TLink(rl))
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(l))
}

func (rRouter *Router) Read(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	l, err := rRouter.hHandler.Read(r.Context(), lid)
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(l))
}

func (rRouter *Router) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}

	rl := TLink{}
	if err = render.Bind(r, &rl); err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	l, err := rRouter.hHandler.Update(r.Context(), lid, handler.TLink(rl))
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(l))
}

func (rRouter *Router) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	l, err := rRouter.hHandler.Delete(r.Context(), lid)
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(l))
}

func (rRouter *Router) Go(w http.ResponseWriter, r *http.Request)
func (rRouter *Router) Stat(w http.ResponseWriter, r *http.Request)
