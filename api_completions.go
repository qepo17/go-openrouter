package openrouter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenRouter normalizes the schema across models and providers
// For most use cases, could be completed by this API
// https://openrouter.ai/docs#request

type ResponseFormat string

var (
	ResponseFormatJSON ResponseFormat = "json_object"
	ResponseFormatText ResponseFormat = "text"
)

type NonChatChoice struct {
	FinishReason *string `json:"finish_reason"` // Nullable
	Text         string  `json:"text"`
}

type CompletionsRequest struct {
	// Prompt is the input text to generate completions from
	// Required
	Prompt string `json:"prompt"`

	// Model is the model to use for the request
	// See list of models here:
	// https://openrouter.ai/docs#models
	// Optional
	// If not provided, the default model set on your openrouter account will be used
	Model *string `json:"model,omitempty"`

	// ResponseFormat is the format of the response
	// Can be used for OpenAI models only
	// Optional
	ResponseFormat *struct{ Type ResponseFormat } `json:"response_format,omitempty"`

	// Seed is the random seed to use for the request
	// Can be used for OpenAI models only
	// Optional
	Seed *int `json:"seed,omitempty"`

	// Stop is the stop sequence for the request
	// Optional
	Stop *[]string `json:"stop,omitempty"`

	// Below are the LLM Parameters can be used
	// For more details, see https://openrouter.ai/docs#llm-parameters
	MaxToken          *int     `json:"max_token,omitempty"`          // 1 - Context Length
	Temperature       *float32 `json:"temperature,omitempty"`        // 0.0 - 2.0
	TopP              *float32 `json:"top_p,omitempty"`              // 0.0 - 1.0
	TopK              *int     `json:"top_k,omitempty"`              // 0 - Infinity. Not available for OpenAI models
	FrequencyPenalty  *float32 `json:"frequency_penalty,omitempty"`  // -2 - 2
	PresencePenalty   *float32 `json:"presence_penalty,omitempty"`   // -2 - 2
	RepetitionPenalty *float32 `json:"repetition_penalty,omitempty"` // 0.0 - 2.0

	// OpenRouter parameters

	// Transform
	// For more detail, see https://openrouter.ai/docs#transforms
	Transforms *[]string `json:"transforms,omitempty"`

	// FallbackModels
	// For more detail, see https://openrouter.ai/docs#oauth
	FallbackModels *[]string `json:"models,omitempty"`
}

func (c CompletionsRequest) validate() error {
	if c.Prompt == "" {
		return ErrPromptIsRequired
	}

	if (c.ResponseFormat != nil && c.Seed != nil) && c.Model != nil && !strings.HasPrefix(*c.Model, "openai") {
		return ErrParameterOnlyForOpenAIModels
	}

	if c.MaxToken != nil && *c.MaxToken < 1 {
		return ErrMaxTokenOutOfRange
	}

	if c.Temperature != nil && (*c.Temperature < 0.0 || *c.Temperature > 2.0) {
		return ErrTemperatureOutOfRange
	}

	if c.TopP != nil && (*c.TopP < 0.0 || *c.TopP > 1.0) {
		return ErrTopPOutOfRange
	}

	if c.TopK != nil && *c.TopK < 0 {
		return ErrTopKOutOfRange
	}

	if c.FrequencyPenalty != nil && (*c.FrequencyPenalty < -2 || *c.FrequencyPenalty > 2) {
		return ErrFrequencyPenaltyOutOfRange
	}

	if c.PresencePenalty != nil && (*c.PresencePenalty < -2 || *c.PresencePenalty > 2) {
		return ErrPresencePenaltyOutOfRange
	}

	if c.RepetitionPenalty != nil && (*c.RepetitionPenalty < 0.0 || *c.RepetitionPenalty > 2.0) {
		return ErrRepetitionPenaltyOutOfRange
	}

	return nil
}

type CompletionsResponse struct {
	ID string `json:"id"`

	// Depending on whether you set "stream" to "true" and
	// whether you passed in "messages" or a "prompt", you
	// will get a different output shape
	// OpenRouter provide four different response formats
	// NonStreamingChoice | StreamingChoice | NonChatChoice | Error
	// But as of now, we not support stream or message yet, so we only support NonChatChoice and Error
	// See https://openrouter.ai/docs#response
	Choices []NonChatChoice `json:"choices"`

	// Unix timestamp of when the request was created
	Created int64 `json:"created"`

	// Model used for the request
	Model string `json:"model"`
}

func (c *Client) Completions(ctx context.Context, req CompletionsRequest) (resp CompletionsResponse, err error) {
	if err := req.validate(); err != nil {
		return resp, err
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	reqURL := c.baseURL + "/api/v1/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", reqURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return resp, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	if c.appURL != "" {
		httpReq.Header.Set("HTTP-Referer", c.appURL)
	}

	if c.appName != "" {
		httpReq.Header.Set("X-Title", c.appName)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("Http Status Code: %d", httpResp.StatusCode)
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
