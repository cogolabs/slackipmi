package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
)

func TestOAuthInvalid(t *testing.T) {
	testflight.WithServer(http.DefaultServeMux, func(r *testflight.Requester) {

		request, err := http.NewRequest("POST", "/oauth", strings.NewReader(""))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		assert.Nil(t, err)
		response := r.Do(request)

		assert.Equal(t, 500, response.StatusCode)
		assert.Equal(t, "invalid_client_id", response.Body)

	})
}
