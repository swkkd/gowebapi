package gowebapi_test

import (
	"context"
	"gowebapi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStreamGetValue(t *testing.T) {
	// mock http server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check header
		username, password, ok := r.BasicAuth()
		if !ok || username != "testuser" || password != "testpass" {
			t.Errorf("auth error")
		}
		// send mock response
		w.Write([]byte(`{"value": "test data"}`))
	}))
	defer server.Close()

	// client setup
	ctx := context.WithValue(context.Background(), gowebapi.AuthContextKey{}, gowebapi.AuthCredentials{Username: "testuser", Password: "testpass"})
	cfg := &gowebapi.Configuration{BasePath: server.URL, HTTPClient: server.Client()}
	client := gowebapi.NewAPIClient(cfg, ctx)

	// function call

	_, _, err := client.StreamApi.StreamGetValue("TestWebId", nil)
	if err != nil {
		t.Errorf("Error calling StreamGetValue: %v", err)
	}
	// add more tests
}
