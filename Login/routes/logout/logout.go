package logout

import (
	"net/http"
	"net/url"
	"os"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("TV_DOMAIN")

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:3000")
	parameters.Add("client_id", os.Getenv("TV_CLIENT_ID"))
	Url.RawQuery = parameters.Encode()

	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}
