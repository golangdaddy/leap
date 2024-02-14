package common

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
)

const (
	CONST_VERTEX_MODEL = "gemini-pro"
)

func (self *GCPClients) GenerateContent(prompt string, temperature float32) (int, *genai.GenerateContentResponse, error) {

	model := self.VertexModel(CONST_VERTEX_MODEL, temperature)

	countResp, err := model.CountTokens(context.Background(), genai.Text(prompt))
	if err != nil {
		return 0, nil, err
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text("What is the average size of a swallow?"))
	if err != nil {
		return 0, nil, err
	}

	return int(countResp.TotalTokens), resp, nil
}
