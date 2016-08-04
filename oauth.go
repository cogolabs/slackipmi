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
	// b, _ := json.Marshal(r.Form)
	// log.Println(string(b))

	token, scope, err := slack.GetOAuthToken(*oauthClientID, *oauthClientSecret, code, *baseURL+"/oauth", false)
	log.Printf("oauth: %+v / %+v / %s\n", token, scope, err)
	if err != nil {
		w.WriteHeader(500)
	}
}
