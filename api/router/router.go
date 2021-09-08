package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sanya-spb/URL-shortener/api/handler"
)

type Router struct {
	http.Handler
	hHandler *handler.Handler
}

type TLink handler.TLink

func (TLink) Bind(r *http.Request) error {
	return nil
}
func (TLink) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewRouter(hHandler *handler.Handler) *Router {
	rRouter := &Router{
		hHandler: hHandler,
	}

	r := chi.NewRouter()
	r.Group(func(rAdm chi.Router) {
		rAdm.Get("/stat", rRouter.Stat)
		rAdm.Put("/{uuid}", rRouter.Update)
		rAdm.Post("/r/{uuid}", rRouter.UpdateRet)
		rAdm.Delete("/delete/{uuid}", rRouter.Delete)
		rAdm.Delete("/delete/r/{uuid}", rRouter.DeleteRet)
		rAdm.Get("/status/{code}", rRouter.StatusCode)
	}) //.Use(auth.AuthMiddleware) // TODO: заменить на middleware.BasicAuth()
	r.Put("/", rRouter.Create)
	r.Get("/info/{uuid}", rRouter.Read)
	r.Get("/status", rRouter.Status)
	r.Get("/", rRouter.Go)

	rRouter.Handler = r
	return rRouter
}

func (rRouter *Router) Create(w http.ResponseWriter, r *http.Request) {
	link := TLink{}
	if err := render.Bind(r, &link); err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.Create(r.Context(), handler.TLink(link))
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, TLink(hLink))
}

func (rRouter *Router) Read(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.Read(r.Context(), lid)
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(hLink))
}

func (rRouter *Router) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}

	link := TLink{}
	if err = render.Bind(r, &link); err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	if err := rRouter.hHandler.Update(r.Context(), lid, handler.TLink(link)); err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}

func (rRouter *Router) UpdateRet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}

	link := TLink{}
	if err = render.Bind(r, &link); err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.UpdateRet(r.Context(), lid, handler.TLink(link))
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(hLink))
}

func (rRouter *Router) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	if err := rRouter.hHandler.Delete(r.Context(), lid); err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Status(r, http.StatusNoContent)
}

func (rRouter *Router) DeleteRet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	lid, err := uuid.Parse(id)
	if err != nil {
		render.Render(w, r, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.DeleteRet(r.Context(), lid)
	if err != nil {
		render.Render(w, r, Err500(err))
		return
	}
	render.Render(w, r, TLink(hLink))
}

func (rRouter *Router) Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (rRouter *Router) StatusCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	fmt.Fprintln(w, code)
}

func (rRouter *Router) Go(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, Err501(errors.New("/{id} not implemented")))
}

func (rRouter *Router) Stat(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, Err501(errors.New("/stat not implemented")))
}
