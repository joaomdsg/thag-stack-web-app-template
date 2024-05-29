package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/joaomdsg/thag-stack-web-app-template/platform/router"
)

//go:embed web/public
var content embed.FS

func main() {
	publicDirContent, err := fs.Sub(content, "web")
	if err != nil {
		log.Fatal(err)
	}
	r := router.New(publicDirContent)
	http.ListenAndServe(":3000", r)
}
