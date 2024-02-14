package openrouter

import "errors"

var (
	ErrApiKeyIsRequired             = errors.New("api key is required")
	ErrPromptIsRequired             = errors.New("prompt is required")
	ErrParameterOnlyForOpenAIModels = errors.New("response format or seed parameter is only available for OpenAI models")
	ErrMaxTokenOutOfRange           = errors.New("max token should be between 1 and context length")
	ErrTemperatureOutOfRange        = errors.New("temperature should be between 0.0 and 2.0")
	ErrTopPOutOfRange               = errors.New("top p should be between 0.0 and 1.0")
	ErrTopKOutOfRange               = errors.New("top k should be between 0 and infinity")
	ErrFrequencyPenaltyOutOfRange   = errors.New("frequency penalty should be between -2 and 2")
	ErrPresencePenaltyOutOfRange    = errors.New("presence penalty should be between -2 and 2")
	ErrRepetitionPenaltyOutOfRange  = errors.New("repetition penalty should be between 0.0 and 2.0")
)
