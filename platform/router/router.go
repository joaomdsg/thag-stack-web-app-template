package router

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaomdsg/thag-stack-web-app-template/platform/handlers"
	"github.com/joaomdsg/thag-stack-web-app-template/web/pages/home"
)

// New returns a chi mux pointer with all the application routes and middleware already configured. It requires
// access to an embedded file system containing a directory named 'public' where static files are stored.
func New(saticContent fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/", templ.Handler(home.Page()))
	r.Mount("/public/", http.FileServer(http.FS(saticContent)))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, saticContent, "public/favicon.ico")
	})

	r.Post("/login", handlers.LoginHandler())

	return r
}
