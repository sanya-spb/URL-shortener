package handler

import (
	"github.com/sanya-spb/URL-shortener/app/repos/links"

	"net/http"
)

type Router struct {
	*http.ServeMux
	links *links.Links
}

func NewRouter(links *links.Links) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		links:    links,
	}
	// r.HandleFunc("/create", r.AuthMiddleware(http.HandlerFunc(r.CreateUser)).ServeHTTP)
	// r.HandleFunc("/read", r.AuthMiddleware(http.HandlerFunc(r.ReadUser)).ServeHTTP)
	// r.HandleFunc("/delete", r.AuthMiddleware(http.HandlerFunc(r.DeleteUser)).ServeHTTP)
	// r.HandleFunc("/search", r.AuthMiddleware(http.HandlerFunc(r.SearchUser)).ServeHTTP)
	return r
}

// TODO: будем брать из пакета links
// type TLinks struct {
// 	ID        uuid.UUID `json:"id"`
// 	URL       string    `json:"url"`
// 	Name      string    `json:"name"`
// 	Descr     string    `json:"descr"`
// 	ShortLink string    `json:"short_link"`
// 	CreatedAt time.Time `json:"created_at"`
// 	DeleteAt  time.Time `json:"delete_at"`
// 	User      string    `json:"user"`
// }
