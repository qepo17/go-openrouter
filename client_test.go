package openrouter

import (
	"net/http"
	"testing"

	"github.com/qepo17/go-openrouter/internal/test"
)

func TestNew(t *testing.T) {
	// Test cases
	tests := []struct {
		name   string
		apiKey string
		want   *Client
		err    error
	}{
		{
			name:   "Valid API key",
			apiKey: "valid-api-key",
			want: &Client{
				baseURL:    baseURL,
				apiKey:     "valid-api-key",
				httpClient: http.DefaultClient,
			},
			err: nil,
		},
		{
			name:   "Empty API key",
			apiKey: "",
			want:   nil,
			err:    ErrApiKeyIsRequired,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.apiKey, ClientOptions{})

			test.Equal(t, tt.want, got, "New()")

			test.ErrorEqual(t, tt.err, err, "New()")
		})
	}
}
