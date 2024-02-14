package openrouter_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mockHttpClient *http.Client
var baseUrl string

func TestMain(m *testing.M) {
	mockServer := serverTest()

	mockHttpClient = mockServer.Client()
	baseUrl = mockServer.URL

	exitCode := m.Run()

	mockServer.Close()

	os.Exit(exitCode)
}

func serverTest() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/chat/completions", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"choices": [{"text": "2"}]}`)
	}))

	return httptest.NewServer(mux)
}
