package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
)

func init() {
	argv0 = "echo"
}

const (
	testActionsPayload  = `payload=%7B%22actions%22%3A%5B%7B%22name%22%3A%22power%22%2C%22type%22%3A%22button%22%2C%22value%22%3A%22reset%22%7D%5D%2C%22callback_id%22%3A%22baremetal1%7CADMIN%7Cfoobar123%22%2C%22team%22%3A%7B%22id%22%3A%22T029F9CR3%22%2C%22domain%22%3A%22myslackorg%22%7D%2C%22channel%22%3A%7B%22id%22%3A%22C12VDUEQH%22%2C%22name%22%3A%22it%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U9ZNJJ9JB%22%2C%22name%22%3A%22joe%22%7D%2C%22action_ts%22%3A%221496393082.299856%22%2C%22message_ts%22%3A%221496393076.318680%22%2C%22attachment_id%22%3A%222%22%2C%22token%22%3A%22V4hafFbeT1doasdfkXeE4f%22%2C%22is_app_unfurl%22%3Afalse%2C%22response_url%22%3A%22https%3A%5C%2F%5C%2Fhooks.slack.com%5C%2Factions%5C%2FT029F9CR3%5C%2F192526491447%5C%2Fxx1yy2zz3%22%7D`
	testActionsExpected = "-H baremetal1 -U ADMIN -P foobar123 power reset\n"
)

var testActions = [][]string{
	{testActionsPayload, testActionsExpected},
	{strings.Replace(testActionsPayload, *slackTeam, "badorg", -1), ""},
	{strings.Replace(testActionsPayload, *slackToken, "badtoken", -1), ""},
	{"", ""},
}

func TestActions(t *testing.T) {
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {
		for _, test := range testActions {

			request, err := http.NewRequest("POST", "/actions", strings.NewReader(test[0]))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			assert.Nil(t, err)
			response := r.Do(request)

			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, test[1], response.Body)

		}
	})
}

func TestActionsMissing(t *testing.T) {
	oldArgv0 := argv0
	argv0 = "missing"
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {

		request, err := http.NewRequest("POST", "/actions", strings.NewReader(testActionsPayload))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		response := r.Do(request)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, `exec: "missing": executable file not found in $PATH`, response.Body)

	})
	argv0 = oldArgv0
}
