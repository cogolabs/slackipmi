package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr    = flag.String("http", ":80", "")
	baseURL = flag.String("base-url", "https://slack.colofoo.net", "base URL (Slack-accessible)")
)

func init() {
	flag.Parse()

	http.Handle("/actions", http.HandlerFunc(actions))
	http.Handle("/oauth", http.HandlerFunc(oauth))
	http.Handle("/power", http.HandlerFunc(power))
}

func main() {
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Println(err)
	}
}
