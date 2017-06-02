package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
)

const (
	testPowerPayload  = `token=V4hafFbeT1doasdfkXeE4f&team_id=T029F9CR3&team_domain=myslackorg&channel_id=C12QQUEQH&channel_name=it&user_id=U9ZNGM9JJ&user_name=joe&command=%2Fpower&text=baremetal1+ADMIN+foobar123&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT029F9CR3%2F191797537237%2Fxx1yy2zz3`
	testPowerExpected = `{
  "text": "baremetal1",
  "attachments": [
      { "text": "-H baremetal1 -U ADMIN -P foobar123 power status" },
      {
          "fallback": "Your slack client does not support interactivity :(",
          "callback_id": "baremetal1|ADMIN|foobar123",
          "color": "#0f0",
          "actions": [
            { "name": "power", "text": "Status", "type": "button", "value": "status" },
            { "name": "power", "text": "On", "type": "button", "value": "on", "style": "primary" }
          ]
      },
      {
          "fallback": "Your slack client does not support interactivity :(",
          "callback_id": "baremetal1|ADMIN|foobar123",
          "color": "#f00",
          "actions": [
            { "name": "power", "text": "Off", "type": "button", "value": "off", "style": "danger" },
            { "name": "power", "text": "Reset", "type": "button", "value": "reset", "style": "danger" },
            { "name": "power", "text": "Cycle", "type": "button", "value": "cycle", "style": "danger" }
        ]
    }
  ]
}`
)

func TestPower(t *testing.T) {
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {

		request, err := http.NewRequest("POST", "/power", strings.NewReader(testPowerPayload))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		response := r.Do(request)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, testPowerExpected, response.Body)

	})
}

func TestPowerInvalid(t *testing.T) {
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {

		request, err := http.NewRequest("POST", "/power", strings.NewReader(""))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		response := r.Do(request)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "invalid host", response.Body)

	})
}

func TestPowerMissing(t *testing.T) {
	oldArgv0 := argv0
	argv0 = "missing"
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {

		request, err := http.NewRequest("POST", "/power", strings.NewReader(testPowerPayload))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		response := r.Do(request)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, `exec: "missing": executable file not found in $PATH`, response.Body)

	})
	argv0 = oldArgv0
}
