package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaomdsg/thag-stack-web-app-template/web/pages/home"
)

// New returns a chi mux
func New(content embed.FS) *chi.Mux {

	publicDir, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/", templ.Handler(home.Page()))
	r.Mount("/public/", http.FileServer(http.FS(publicDir)))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, publicDir, "public/favicon.ico")
	})

	return r
}
