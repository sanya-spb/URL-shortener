package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (hHandler *Handler) Create(ctx context.Context, link TLink) (TLink, error) {
	data, err := hHandler.links.Create(ctx, links.TLink(link))
	if err != nil {
		return TLink{}, fmt.Errorf("error when creating: %w", err)
	}

	return TLink(*data), nil
}

func (hHandler *Handler) Read(ctx context.Context, id string) (TLink, error) {
	if id == "" {
		return TLink{}, fmt.Errorf("bad request: ID is empty")
	}

	data, err := hHandler.links.Read(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when reading: %w", err)
	}

	return TLink(*data), nil
}

func (hHandler *Handler) Update(ctx context.Context, id string, data TLink) error {
	if id == "" {
		return fmt.Errorf("bad request: ID is empty")
	}

	err := hHandler.links.Update(ctx, id, links.TLink(data))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrLinkNotFound
		}
		return fmt.Errorf("error when updating: %w", err)
	}

	return nil
}

func (hHandler *Handler) UpdateRet(ctx context.Context, id string, data TLink) (TLink, error) {
	if id == "" {
		return TLink{}, fmt.Errorf("bad request: ID is empty")
	}

	newData, err := hHandler.links.UpdateRet(ctx, id, links.TLink(data))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when updating: %w", err)
	}

	return TLink(*newData), nil
}

func (hHandler *Handler) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("bad request: ID is empty")
	}

	err := hHandler.links.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrLinkNotFound
		}
		return fmt.Errorf("error when deleting: %w", err)
	}

	return nil
}

func (hHandler *Handler) DeleteRet(ctx context.Context, id string) (TLink, error) {
	if id == "" {
		return TLink{}, fmt.Errorf("bad request: ID is empty")
	}

	delData, err := hHandler.links.DeleteRet(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TLink{}, ErrLinkNotFound
		}
		return TLink{}, fmt.Errorf("error when deleting: %w", err)
	}

	return TLink(*delData), nil
}

func (hHandler *Handler) Go(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("bad request: ID is empty")
	}

	data, err := hHandler.links.Go(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrLinkNotFound
		}
		return "", fmt.Errorf("error when reading: %w", err)
	}

	return data, nil
}

func (hHandler *Handler) Stat(ctx context.Context) (chan TLink, error) {
	chin, err := hHandler.links.Stat(ctx)
	if err != nil {
		return nil, err
	}
	chout := make(chan TLink, 100)

	go func() {
		defer close(chout)
		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-chin:
				if !ok {
					return
				}
				chout <- TLink(data)
			}
		}
	}()

	return chout, nil
}
