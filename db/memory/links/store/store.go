package store

import (
	"context"
	"database/sql"
	"math/rand"
	"sync"
	"time"

	"github.com/sanya-spb/URL-shortener/app/repos/links"
)

var _ links.LinksStore = &Links{}

type Links struct {
	sync.RWMutex
	m map[string]links.TLink
}

func NewLinks() *Links {
	return &Links{
		m: make(map[string]links.TLink),
	}
}

func (link *Links) Create(ctx context.Context, data links.TLink) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	id, err := link.GetNextID(ctx)
	if err != nil {
		return "", nil
	}
	data.ID = id

	link.Lock()
	defer link.Unlock()

	link.m[data.ID] = data
	return id, nil
}

func (link *Links) Read(ctx context.Context, id string) (*links.TLink, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	link.RLock()
	defer link.RUnlock()

	data, ok := link.m[id]
	if ok {
		return &data, nil
	}
	return nil, sql.ErrNoRows
}

func (link *Links) Update(ctx context.Context, id string, data links.TLink) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if ok, err := link.IsExist(ctx, id); err != nil {
		return err
	} else {
		if !ok {
			return sql.ErrNoRows
		}
	}

	data.ID = id

	link.Lock()
	defer link.Unlock()

	link.m[data.ID] = data
	return nil
}

func (link *Links) UpdateRet(ctx context.Context, id string, data links.TLink) (*links.TLink, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if ok, err := link.IsExist(ctx, id); err != nil {
		return nil, err
	} else {
		if !ok {
			return nil, sql.ErrNoRows
		}
	}

	data.ID = id

	link.Lock()
	defer link.Unlock()

	link.m[data.ID] = data
	return &data, nil
}

func (link *Links) Delete(ctx context.Context, id string) error {
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

func (link *Links) DeleteRet(ctx context.Context, id string) (*links.TLink, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	link.Lock()
	defer link.Unlock()

	data, ok := link.m[id]
	if ok {
		delete(link.m, id)
		return &data, nil
	}
	return nil, sql.ErrNoRows
}

func (link *Links) IsExist(ctx context.Context, id string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	}

	link.RLock()
	defer link.RUnlock()

	_, ok := link.m[id]
	if ok {
		return true, nil
	}
	return false, nil
}

func (link *Links) GetNextID(ctx context.Context) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	n := 6

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		if ok, err := link.IsExist(ctx, string(b)); err != nil {
			return "", err
		} else {
			if ok {
				continue
			}
		}
		return string(b), nil
	}
}

func (link *Links) Go(ctx context.Context, id string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	link.Lock()
	defer link.Unlock()

	data, ok := link.m[id]
	if ok {
		if link.m[id].DeleteAt.Before(time.Now()) {
			return "", sql.ErrNoRows
		}
		data.GoCount++
		link.m[id] = data
		return data.URL, nil
	}
	return "", sql.ErrNoRows
}

func (link *Links) Stat(ctx context.Context) (chan links.TLink, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	chout := make(chan links.TLink, 100)

	go func() {
		defer close(chout)

		link.RLock()
		defer link.RUnlock()

		for _, data := range link.m {
			select {
			case <-ctx.Done():
				return
			case chout <- data:
			}

		}
	}()

	return chout, nil
}
