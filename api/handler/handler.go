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

func (hHandler *Handler) Create(ctx context.Context, link TLink) (TLink, error) {
	data, err := hHandler.links.Create(ctx, links.TLink(link))
	if err != nil {
		return TLink{}, fmt.Errorf("error when creating: %w", err)
	}

	return TLink(*data), nil
}

func (hHandler *Handler) Read(ctx context.Context, id uuid.UUID) (TLink, error) {
	if (id == uuid.UUID{}) {
		return TLink{}, fmt.Errorf("bad request: UUID is empty")
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

func (hHandler *Handler) Update(ctx context.Context, id uuid.UUID, data TLink) error {
	if (id == uuid.UUID{}) {
		return fmt.Errorf("bad request: UUID is empty")
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

func (hHandler *Handler) UpdateRet(ctx context.Context, id uuid.UUID, data TLink) (TLink, error) {
	if (id == uuid.UUID{}) {
		return TLink{}, fmt.Errorf("bad request: UUID is empty")
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

func (hHandler *Handler) Delete(ctx context.Context, id uuid.UUID) error {
	if (id == uuid.UUID{}) {
		return fmt.Errorf("bad request: UUID is empty")
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

func (hHandler *Handler) DeleteRet(ctx context.Context, id uuid.UUID) (TLink, error) {
	if (id == uuid.UUID{}) {
		return TLink{}, fmt.Errorf("bad request: UUID is empty")
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
