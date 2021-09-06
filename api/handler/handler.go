package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sanya-spb/URL-shortener/app/repos/links"
)

type Handler struct {
	links *links.Links
}

func NewHandler(links *links.Links) *Handler {
	r := &Handler{
		links: links,
	}
	// r.HandleFunc("/create", r.AuthMiddleware(http.HandlerFunc(r.CreateUser)).ServeHTTP)
	// r.HandleFunc("/read", r.AuthMiddleware(http.HandlerFunc(r.ReadUser)).ServeHTTP)
	// r.HandleFunc("/delete", r.AuthMiddleware(http.HandlerFunc(r.DeleteUser)).ServeHTTP)
	// r.HandleFunc("/search", r.AuthMiddleware(http.HandlerFunc(r.SearchUser)).ServeHTTP)
	return r
}

// TODO: пока берем из пакета links, потом решим что тут лишнее
// type TLinks struct {
// 	ID        uuid.UUID `json:"id"`
// 	URL       string    `json:"url"`
// 	Name      string    `json:"name"`
// 	Descr     string    `json:"descr"`
// 	ShortLink string    `json:"short_link"`
// 	CreatedAt time.Time `json:"created_at"`
// 	DeleteAt  time.Time `json:"delete_at"`
// 	User      string    `json:"user"`
// }

func (hHandler *Handler) Create(ctx context.Context, l links.TLinks) (links.TLinks, error) {
	link := links.TLinks{
		ID:        l.ID,
		URL:       l.URL,
		Name:      l.Name,
		Descr:     l.Descr,
		ShortLink: l.ShortLink,
		CreatedAt: l.CreatedAt,
		DeleteAt:  l.DeleteAt,
		User:      l.User,
	}

	storeLink, err := hHandler.links.Create(ctx, link)
	if err != nil {
		return links.TLinks{}, fmt.Errorf("error when creating: %w", err)
	}

	return links.TLinks{
		ID:        storeLink.ID,
		URL:       storeLink.URL,
		Name:      storeLink.Name,
		Descr:     storeLink.Descr,
		ShortLink: storeLink.ShortLink,
		CreatedAt: storeLink.CreatedAt,
		DeleteAt:  storeLink.DeleteAt,
		User:      storeLink.User,
	}, nil
}

var ErrLinkNotFound = errors.New("link not found")

func (hHandler *Handler) Read(ctx context.Context, uid uuid.UUID) (links.TLinks, error) {
	if (uid == uuid.UUID{}) {
		return links.TLinks{}, fmt.Errorf("bad request: uid is empty")
	}

	storeLink, err := hHandler.links.Read(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return links.TLinks{}, ErrLinkNotFound
		}
		return links.TLinks{}, fmt.Errorf("error when reading: %w", err)
	}

	return links.TLinks{
		ID:        storeLink.ID,
		URL:       storeLink.URL,
		Name:      storeLink.Name,
		Descr:     storeLink.Descr,
		ShortLink: storeLink.ShortLink,
		CreatedAt: storeLink.CreatedAt,
		DeleteAt:  storeLink.DeleteAt,
		User:      storeLink.User,
	}, nil
}

func (hHandler *Handler) Delete(ctx context.Context, uid uuid.UUID) (links.TLinks, error) {
	if (uid == uuid.UUID{}) {
		return links.TLinks{}, fmt.Errorf("bad request: uid is empty")
	}

	storeLink, err := hHandler.links.Delete(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return links.TLinks{}, ErrLinkNotFound
		}
		return links.TLinks{}, fmt.Errorf("error when reading: %w", err)
	}

	return links.TLinks{
		ID:        storeLink.ID,
		URL:       storeLink.URL,
		Name:      storeLink.Name,
		Descr:     storeLink.Descr,
		ShortLink: storeLink.ShortLink,
		CreatedAt: storeLink.CreatedAt,
		DeleteAt:  storeLink.DeleteAt,
		User:      storeLink.User,
	}, nil
}
