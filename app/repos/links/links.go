package links

import (
	"context"
	"fmt"
	"time"
)

type TLink struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Descr     string    `json:"descr"`
	CreatedAt time.Time `json:"created_at"`
	DeleteAt  time.Time `json:"delete_at"`
	User      string    `json:"user"`
}

type LinksStore interface {
	Create(ctx context.Context, data TLink) (string, error)
	Read(ctx context.Context, id string) (*TLink, error)
	Update(ctx context.Context, id string, data TLink) error
	UpdateRet(ctx context.Context, id string, data TLink) (*TLink, error)
	Delete(ctx context.Context, id string) error
	DeleteRet(ctx context.Context, id string) (*TLink, error)
	IsExist(ctx context.Context, id string) (bool, error)
	GetNextID(ctx context.Context) (string, error)
	Go(ctx context.Context, id string) (string, error)
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
	data.ID = id
	return &data, nil
}

// Return Link by ID
func (link *Links) Read(ctx context.Context, id string) (*TLink, error) {
	data, err := link.store.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("read link error: %w", err)
	}
	return data, nil
}

// Update Link by ID
func (link *Links) Update(ctx context.Context, id string, data TLink) error {
	err := link.store.Update(ctx, id, data)
	if err != nil {
		return fmt.Errorf("update link error: %w", err)
	}
	return nil
}

// Update Link by ID with returning updated Link
func (link *Links) UpdateRet(ctx context.Context, id string, data TLink) (*TLink, error) {
	lNew, err := link.store.UpdateRet(ctx, id, data)
	if err != nil {
		return nil, fmt.Errorf("update link error: %w", err)
	}
	return lNew, nil
}

// Delete Link by ID
func (link *Links) Delete(ctx context.Context, id string) error {
	err := link.store.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete link error: %w", err)
	}
	return nil
}

// Delete by ID with returning deleted Link
func (link *Links) DeleteRet(ctx context.Context, id string) (*TLink, error) {
	dLink, err := link.store.DeleteRet(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete link error: %w", err)
	}
	return dLink, nil
}

//
func (link *Links) Go(ctx context.Context, id string) (string, error) {
	// return url.URL{ Scheme: "https", Host: r.Host, Path: r.URL.Path, RawQuery: r.URL.RawQuery, }
	data, err := link.store.Go(ctx, id)
	if err != nil {
		return "", fmt.Errorf("go link error: %w", err)
	}
	return data, nil
}
