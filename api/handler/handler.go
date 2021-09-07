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
	return r
}

// TODO: пока берем из пакета links, потом решим что тут лишнее
type TLink links.TLink

func (hHandler *Handler) Create(ctx context.Context, l TLink) (TLink, error) {
	link := links.TLink{
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
		return TLink{}, fmt.Errorf("error when creating: %w", err)
	}

	return TLink{
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

func (hHandler *Handler) Read(ctx context.Context, uid uuid.UUID) (TLink, error) {
	if (uid == uuid.UUID{}) {
		return TLink{}, fmt.Errorf("bad request: uid is empty")
	}

	storeLink, err := hHandler.links.Read(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when reading: %w", err)
	}

	return TLink{
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

func (hHandler *Handler) Update(ctx context.Context, uid uuid.UUID, l TLink) (TLink, error) {
	if _, err := hHandler.Delete(ctx, uid); err != nil {
		return TLink{}, err
	}
	link, err := hHandler.Create(ctx, l)
	if err != nil {
		return TLink{}, err
	}
	return link, nil
}

func (hHandler *Handler) Delete(ctx context.Context, uid uuid.UUID) (TLink, error) {
	if (uid == uuid.UUID{}) {
		return TLink{}, fmt.Errorf("bad request: uid is empty")
	}

	storeLink, err := hHandler.links.Delete(ctx, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when reading: %w", err)
	}

	return TLink{
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
