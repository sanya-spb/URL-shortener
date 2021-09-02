package links

import (
	"time"

	"github.com/google/uuid"
)

type TLinks struct {
	ID        uuid.UUID `json:"id"`
	URL       string    `json:"url"`
	Name      string    `json:"name"`
	Descr     string    `json:"descr"`
	ShortLink string    `json:"short_link"`
	CreatedAt time.Time `json:"created_at"`
	DeleteAt  time.Time `json:"delete_at"`
	User      string    `json:"user"`
}

type LinksStore interface {
	// Create(ctx context.Context, u User) (*uuid.UUID, error)
	// Read(ctx context.Context, uid uuid.UUID) (*User, error)
	// Delete(ctx context.Context, uid uuid.UUID) error
	// SearchUsers(ctx context.Context, s string) (chan User, error)
}

type Links struct {
	store LinksStore
}

func NewLinks(store LinksStore) *Links {
	return &Links{
		store: store,
	}
}
