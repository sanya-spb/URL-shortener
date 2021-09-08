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
	Create(ctx context.Context, data TLink) (*uuid.UUID, error)
	Read(ctx context.Context, id uuid.UUID) (*TLink, error)
	Update(ctx context.Context, id uuid.UUID, data TLink) error
	UpdateRet(ctx context.Context, id uuid.UUID, data TLink) (*TLink, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteRet(ctx context.Context, id uuid.UUID) (*TLink, error)
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

// Create new Link with returning it
func (link *Links) Create(ctx context.Context, data TLink) (*TLink, error) {
	id, err := link.store.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("create link error: %w", err)
	}
	data.ID = *id
	return &data, nil
}

// Return Link by UUID
func (link *Links) Read(ctx context.Context, id uuid.UUID) (*TLink, error) {
	data, err := link.store.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return data, nil
}

// Update Link by UUID
func (link *Links) Update(ctx context.Context, id uuid.UUID, data TLink) error {
	err := link.store.Update(ctx, id, data)
	if err != nil {
		return fmt.Errorf("update link error: %w", err)
	}
	return nil
}

// Update Link by UUID with returning updated Link
func (link *Links) UpdateRet(ctx context.Context, id uuid.UUID, data TLink) (*TLink, error) {
	lNew, err := link.store.UpdateRet(ctx, id, data)
	if err != nil {
		return nil, fmt.Errorf("update link error: %w", err)
	}
	return lNew, nil
}

// Delete Link by UUID
func (link *Links) Delete(ctx context.Context, lid uuid.UUID) error {
	err := link.store.Delete(ctx, lid)
	if err != nil {
		return fmt.Errorf("delete link error: %w", err)
	}
	return nil
}

// Delete by UUID with returning deleted Link
func (link *Links) DeleteRet(ctx context.Context, lid uuid.UUID) (*TLink, error) {
	dLink, err := link.store.DeleteRet(ctx, lid)
	if err != nil {
		return nil, fmt.Errorf("delete link error: %w", err)
	}
	return dLink, nil
}
