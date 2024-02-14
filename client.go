package openrouter

import (
	"net/http"
)

type Client struct {
	baseURL string
	apiKey  string

	httpClient *http.Client

	appURL  string
	appName string
}

type ClientOptions struct {
	AppURL  string // Optional, for including your app on openrouter.ai rankings.
	AppName string // Optional. Shows in rankings on openrouter.ai.

	APITimeoutInSeconds int // Optional. Default is 30 seconds.

	HttpClient *http.Client // Optional. Default is http.DefaultClient.

	BaseURL string // Optional. Default is "https://openrouter.ai".
}

func New(apiKey string, opts ClientOptions) (*Client, error) {
	if apiKey == "" {
		return nil, ErrApiKeyIsRequired
	}

	if opts.APITimeoutInSeconds == 0 {
		opts.APITimeoutInSeconds = 30
	}

	if opts.HttpClient == nil {
		opts.HttpClient = http.DefaultClient
	}

	if opts.BaseURL == "" {
		opts.BaseURL = "https://openrouter.ai"
	}

	return &Client{
		baseURL:    opts.BaseURL,
		apiKey:     apiKey,
		appURL:     opts.AppURL,
		appName:    opts.AppName,
		httpClient: opts.HttpClient,
	}, nil
}
