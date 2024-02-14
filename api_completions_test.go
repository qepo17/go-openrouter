package openrouter_test

import (
	"context"
	"testing"

	"github.com/qepo17/go-openrouter/internal/test"

	"github.com/qepo17/go-openrouter"
)

func TestCompletions(t *testing.T) {
	// Test cases
	tests := []struct {
		name string
		req  openrouter.CompletionsRequest
		want openrouter.CompletionsResponse
		err  error
	}{
		{
			name: "empty prompt, should return error",
			req:  openrouter.CompletionsRequest{},
			want: openrouter.CompletionsResponse{},
			err:  openrouter.ErrPromptIsRequired,
		},
		{
			name: "valid request",
			req: openrouter.CompletionsRequest{
				Prompt: "Answer without explanation. What is 1 + 1 =",
			},
			want: openrouter.CompletionsResponse{
				Choices: []openrouter.NonChatChoice{
					{
						Text: "2",
					},
				},
			},
			err: nil,
		},
	}

	client, err := openrouter.New("valid-api-key", openrouter.ClientOptions{
		HttpClient: mockHttpClient,
		BaseURL:    baseUrl,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Run tests
	for _, tt := range tests {
		got, err := client.Completions(context.Background(), tt.req)
		test.Equal(t, tt.want, got, "Completions()")
		test.ErrorEqual(t, tt.err, err, "Completions()")
	}
}
