package handlers

import (
	"log"
	"net/http"
)

type httpHandlerFuncThatErrors func(w http.ResponseWriter, r *http.Request) error

func withErrorHandling(handler httpHandlerFuncThatErrors) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			// handle error
			log.Print(err)

		}
	}
}
