package links

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TLink struct {
	ID        uuid.UUID `json:"id"`
	URL       string    `json:"url"`
	Name      string    `json:"name"`
	Descr     string    `json:"descr"`
	LinkID    string    `json:"link_id"`
	CreatedAt time.Time `json:"created_at"`
	DeleteAt  time.Time `json:"delete_at"`
	User      string    `json:"user"`
}

type LinksStore interface {
	Create(ctx context.Context, l TLink) (*uuid.UUID, error)
	Read(ctx context.Context, lid uuid.UUID) (*TLink, error)
	// Update(ctx context.Context, lid uuid.UUID, l TLinks) (*uuid.UUID, error)
	Delete(ctx context.Context, lid uuid.UUID) error
	// Go(ctx context.Context, sl string) error
	// Stat(ctx context.Context, s string) (chan TLinks, error)
}

type Links struct {
	store LinksStore
}

func NewLinks(store LinksStore) *Links {
	return &Links{
		store: store,
	}
}

func (link *Links) Create(ctx context.Context, l TLink) (*TLink, error) {
	id, err := link.store.Create(ctx, l)
	if err != nil {
		return nil, fmt.Errorf("create link error: %w", err)
	}
	l.ID = *id
	return &l, nil
}

func (link *Links) Read(ctx context.Context, lid uuid.UUID) (*TLink, error) {
	l, err := link.store.Read(ctx, lid)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return l, nil
}

func (link *Links) Update(ctx context.Context, lid uuid.UUID, l TLink) (*TLink, error) {
	if _, err := link.Delete(ctx, lid); err != nil {
		return nil, err
	}
	lNew, err := link.Create(ctx, l)
	if err != nil {
		return nil, err
	}
	return lNew, nil
}

func (link *Links) Delete(ctx context.Context, lid uuid.UUID) (*TLink, error) {
	l, err := link.store.Read(ctx, lid)
	if err != nil {
		return nil, fmt.Errorf("delete link error: %w", err)
	}
	return l, link.store.Delete(ctx, lid)
}
