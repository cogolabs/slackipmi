package main

import (
	"net/http"
	"strings"
)

func power(w http.ResponseWriter, r *http.Request) {
	text := strings.Split(r.FormValue("text"), " ")
	host := text[0]

	if host == "" {
		w.Write([]byte("invalid host"))
		return
	}

	status, err := ipmipower(host, text[1], text[2], "status")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	powerT.Execute(w, map[string]string{
		"title":    host,
		"status":   strings.TrimSpace(string(status)),
		"callback": host + "|" + strings.Join(text[1:], "|"),
	})
}
