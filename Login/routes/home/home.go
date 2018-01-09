package home

import (
	"html/template"
	"net/http"
	"os"

	templates ".."
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		TVClientId     string
		TVClientSecret string
		TVDomain       string
		TVCallbackURL  template.URL
	}{
		os.Getenv("TV_CLIENT_ID"),
		os.Getenv("TV_CLIENT_SECRET"),
		os.Getenv("TV_DOMAIN"),
		template.URL(os.Getenv("TV_CALLBACK_URL")),
	}

	templates.RenderTemplate(w, "home", data)
}
