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
	m map[uuid.UUID]links.TLinks
}

func NewLinks() *Links {
	return &Links{
		m: make(map[uuid.UUID]links.TLinks),
	}
}

func (link *Links) Create(ctx context.Context, l links.TLinks) (*uuid.UUID, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	uid := uuid.New()
	l.ID = uid
	link.m[l.ID] = l
	return &uid, nil
}

func (link *Links) Read(ctx context.Context, uid uuid.UUID) (*links.TLinks, error) {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	u, ok := link.m[uid]
	if ok {
		return &u, nil
	}
	return nil, sql.ErrNoRows
}

// TODO: вернуть ошибку если не нашли
func (link *Links) Delete(ctx context.Context, uid uuid.UUID) error {
	link.Lock()
	defer link.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	delete(link.m, uid)
	return nil
}
