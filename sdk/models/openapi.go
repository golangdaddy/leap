package models

type ChatCompletionRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}

type GPT3Request struct {
	Model       string                `json:"model"`
	Messages    []*GPT3RequestMessage `json:"messages"`
	Temperature float64               `json:"temperature"`
}

type GPT3RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
