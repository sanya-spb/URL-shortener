package router

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Group(func(rAdm chi.Router) {
		rAdm.Get("/stat", rRouter.Stat)
		rAdm.Put("/{id}", rRouter.Update)
		rAdm.Post("/r/{id}", rRouter.UpdateRet)
		rAdm.Delete("/d/{id}", rRouter.Delete)
		rAdm.Delete("/d/r/{id}", rRouter.DeleteRet)
		rAdm.Get("/status/{code}", rRouter.StatusCode)
	})
	r.Put("/", rRouter.Create)
	r.Get("/i/{id}", rRouter.Read)
	r.Get("/status", rRouter.Status)
	r.Get("/{id}", rRouter.Go)

	r.Get("/ui/*", rRouter.ui)

	rRouter.Handler = r
	return rRouter
}

func (rRouter *Router) Create(w http.ResponseWriter, req *http.Request) {
	link := TLink{}
	if err := render.Bind(req, &link); err != nil {
		render.Render(w, req, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.Create(req.Context(), handler.TLink(link))
	if err != nil {
		render.Render(w, req, Err500(err))
		return
	}
	render.Status(req, http.StatusCreated)
	render.Render(w, req, TLink(hLink))
}

func (rRouter *Router) Read(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	hLink, err := rRouter.hHandler.Read(req.Context(), id)
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Render(w, req, TLink(hLink))
}

func (rRouter *Router) Update(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	link := TLink{}
	if err := render.Bind(req, &link); err != nil {
		render.Render(w, req, Err400(err))
		return
	}
	if err := rRouter.hHandler.Update(req.Context(), id, handler.TLink(link)); err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Status(req, http.StatusNoContent)
}

func (rRouter *Router) UpdateRet(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	link := TLink{}
	if err := render.Bind(req, &link); err != nil {
		render.Render(w, req, Err400(err))
		return
	}
	hLink, err := rRouter.hHandler.UpdateRet(req.Context(), id, handler.TLink(link))
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Render(w, req, TLink(hLink))
}

func (rRouter *Router) Delete(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if err := rRouter.hHandler.Delete(req.Context(), id); err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Status(req, http.StatusNoContent)
}

func (rRouter *Router) DeleteRet(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	hLink, err := rRouter.hHandler.DeleteRet(req.Context(), id)
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}
	render.Render(w, req, TLink(hLink))
}

func (rRouter *Router) Status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (rRouter *Router) StatusCode(w http.ResponseWriter, req *http.Request) {
	code := chi.URLParam(req, "code")
	fmt.Fprintln(w, code)
}

func (rRouter *Router) Go(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	url, err := rRouter.hHandler.Go(req.Context(), id)
	if err != nil {
		if errors.As(err, &handler.ErrLinkNotFound) {
			render.Render(w, req, Err404(err))
			return
		}
		render.Render(w, req, Err500(err))
		return
	}

	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (rRouter *Router) Stat(w http.ResponseWriter, req *http.Request) {
	render.Render(w, req, Err501(errors.New("/stat not implemented")))
}

func (rRouter *Router) ui(w http.ResponseWriter, req *http.Request) {
	root := "./data/ui"
	fs := http.FileServer(http.Dir(root))
	if _, err := os.Stat(root + req.RequestURI); os.IsNotExist(err) {
		http.StripPrefix(req.RequestURI, fs).ServeHTTP(w, req)
	} else {
		fs.ServeHTTP(w, req)
	}
}
