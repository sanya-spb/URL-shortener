package store

import (
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
