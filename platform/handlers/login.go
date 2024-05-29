package handlers

import (
	"context"
	"net/http"

	"github.com/joaomdsg/thag-stack-web-app-template/web/components"
)

func LoginHandler() http.HandlerFunc {
	return withErrorHandling(func(w http.ResponseWriter, r *http.Request) error {
		email := r.FormValue("email")
		password := r.FormValue("password")
		model := components.LoginFormModel{
			ServerErr:   "Oops! Something went wrong.",
			EmailVal:    email,
			PasswordVal: password,
		}
		if email == "" {
			model.EmailErr = "Please enter your email address."
		}
		if password == "" {
			model.PasswordErr = "Please enter your password."
		}

		loginform := components.Loginform(model)
		err := loginform.Render(context.TODO(), w)
		return err
	})
}
