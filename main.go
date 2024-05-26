package main

import (
	"embed"
	"net/http"

	"github.com/joaomdsg/thag-stack-web-app-template/platform/router"
)

//go:embed web/public
var content embed.FS

func main() {
	r := router.New(content)
	http.ListenAndServe(":3000", r)
}
