package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaomdsg/thag-stack-web-app-template/web/pages/home"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/", templ.Handler(home.Page()))
	r.Mount("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("web/public"))))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/public/favicon.ico")
	})

	r.Get("/sse", handleSSE)
	return r
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Send the "server started" event
	fmt.Fprint(w, "event: server-started\ndata: Server started\n\n", time.Now().String())
	w.(http.Flusher).Flush()

	// Wait for the connection to be closed
	<-r.Context().Done()
}
