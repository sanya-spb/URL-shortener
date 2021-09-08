package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/google/uuid"
	"github.com/sanya-spb/URL-shortener/app/repos/links"
)

var _ links.LinksStore = &Links{}

type Links struct {
	sync.Mutex
	m map[uuid.UUID]links.TLink
}

func NewLinks() *Links {
	return &Links{
		m: make(map[uuid.UUID]links.TLink),
	}
}

func (link *Links) Create(ctx context.Context, data links.TLink) (*uuid.UUID, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	uid := uuid.New()
	data.ID = uid
	link.m[data.ID] = data
	return &uid, nil
}

func (link *Links) Read(ctx context.Context, id uuid.UUID) (*links.TLink, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	data, ok := link.m[id]
	if ok {
		return &data, nil
	}
	return nil, sql.ErrNoRows
}

func (link *Links) Update(ctx context.Context, id uuid.UUID, data links.TLink) error {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := link.m[id]
	if !ok {
		return sql.ErrNoRows
	}

	data.ID = id
	link.m[data.ID] = data
	return nil
}

func (link *Links) UpdateRet(ctx context.Context, id uuid.UUID, data links.TLink) (*links.TLink, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	_, ok := link.m[id]
	if !ok {
		return nil, sql.ErrNoRows
	}

	data.ID = id
	link.m[data.ID] = data
	return &data, nil
}

func (link *Links) Delete(ctx context.Context, id uuid.UUID) error {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := link.m[id]
	if ok {
		delete(link.m, id)
		return nil
	}
	return sql.ErrNoRows
}

func (link *Links) DeleteRet(ctx context.Context, id uuid.UUID) (*links.TLink, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	l, ok := link.m[id]
	if ok {
		delete(link.m, id)
		return &l, nil
	}
	return nil, sql.ErrNoRows
}
