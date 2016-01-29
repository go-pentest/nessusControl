package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostWithArgs(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		unmarshalledJSON := `{"token": "TestToken"}`
		fmt.Fprintln(w, unmarshalledJSON)
	}))
	defer testServer.Close()

	jsonStr, err := json.Marshal(`{"username": "test", "password": "test"}`)
	if err != nil {
		t.FailNow()
	}

	client := Client{}

	// Ignore bad HTTPS certificate
	transportSettings := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: transportSettings}

	body, err := client.postWithJSON(httpClient, testServer.URL, jsonStr)
	if err != nil {
		t.FailNow()
	}

	var session createSession
	json.Unmarshal(body, &session)
	if session.Token != "TestToken" {
		t.FailNow()
	}
}
