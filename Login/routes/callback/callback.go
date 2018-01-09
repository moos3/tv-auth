package callback

import (
	"context"
	_ "crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"../../app"
	"golang.org/x/oauth2"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("TV_DOMAIN")

	conf := &oauth2.Config{
		ClientID:     os.Getenv("TV_CLIENT_ID"),
		ClientSecret: os.Getenv("TV_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("TV_CALLBACK_URL"),
		//Scopes:       []string{"*"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/login/oauth/authorize",
			TokenURL: "https://" + domain + "/login/oauth/access_token/",
		},
	}

	state := r.URL.Query().Get("state")
	session, err := app.Store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if state != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusInternalServerError)
		return
	}

	code := r.URL.Query().Get("code")

	token, err := conf.Exchange(context.TODO(), code)
	fmt.Printf("%+v\n", token)
	fmt.Printf("%+v\n", "Code: "+string(code)+"\n")
	fmt.Printf("%+v\n", "State: "+string(state)+"\n")
	os.Exit(3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	client := conf.Client(context.TODO(), token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	print(resp)

	var profile map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err = app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = token.Extra("id_token")
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to logged in page
	http.Redirect(w, r, "/user", http.StatusSeeOther)

}
