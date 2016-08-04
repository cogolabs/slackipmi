package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	slackTeam  = flag.String("slack-team", "myslackorg", "ignore requests outside your Slack Team")
	slackToken = flag.String("slack-token", "V4hafFbeT1doasdfkXeE4f", "token verifies reqs are actually coming from Slack")
)

func actions(w http.ResponseWriter, r *http.Request) {
	payload := r.FormValue("payload")
	action := &Action{}
	err := json.Unmarshal([]byte(payload), &action)
	if err != nil {
		log.Println(err)
		return
	}
	if action.Team.Domain != *slackTeam {
		return
	}
	if action.Token != *slackToken {
		return
	}

	text := strings.Split(action.CallbackID, "|")
	b, err := ipmipower(text[0], text[1], text[2], action.Actions[0].Value)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}
