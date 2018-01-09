package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"

	"../../app"
	"golang.org/x/oauth2"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("TV_DOMAIN")
	aud := os.Getenv("TV_AUDIENCE")

	conf := &oauth2.Config{
		ClientID:     os.Getenv("TV_CLIENT_ID"),
		ClientSecret: os.Getenv("TV_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("TV_CALLBACK_URL"),
		//Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/login/oauth/authorize",
			TokenURL: "https://" + domain + "/login/oauth/access_token",
		},
	}

	if aud == "" {
		aud = "https://" + domain + "/users/me"
	}

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	session, err := app.Store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
