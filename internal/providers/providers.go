package providers

import (
	"context"

	openai "github.com/sashabaranov/go-openai"

	"github.com/numoru-ia/geo-audit/internal/config"
)

type Registry struct {
	client *openai.Client
}

func New(baseURL, apiKey string) *Registry {
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL
	return &Registry{client: openai.NewClientWithConfig(cfg)}
}

func (r *Registry) Ask(ctx context.Context, p config.Provider, query string) (string, error) {
	resp, err := r.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       p.Model,
		Temperature: 0.0,
		MaxTokens:   800,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Responde breve, con fuentes si aplica."},
			{Role: "user", Content: query},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
