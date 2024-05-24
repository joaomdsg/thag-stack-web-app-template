package main

import (
	"net/http"

	"github.com/joaomdsg/thag-stack-web-app-template/platform/router"
)

func main() {
	r := router.New()
	http.ListenAndServe(":3000", r)
}
