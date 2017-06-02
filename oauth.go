package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"
)

var (
	oauthClientID     = flag.String("oauth-client-id", os.Getenv("OAUTH_CLIENT_ID"), "Slack-provided client ID")
	oauthClientSecret = flag.String("oauth-client-secret", os.Getenv("OAUTH_CLIENT_SECRET"), "Slack-provided client secret")
)

func oauth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	code := r.FormValue("code")

	token, scope, err := slack.GetOAuthToken(*oauthClientID, *oauthClientSecret, code, *baseURL+"/oauth", false)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("oauth: %+v / %+v\n", token, scope)
}
