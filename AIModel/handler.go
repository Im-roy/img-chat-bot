package aimodel

import (
	"context"
	"img-chat-bot/AIModel/gemini"
	"img-chat-bot/model"
)

type AiModel struct {
	AIClient gemini.GeminiAI
}

func (ai *AiModel) GenerateResponse(ctx context.Context, promptMessage string, additionalImageData []model.PromptImageModel) (string, error) {
	resp, err := ai.AIClient.GenerateContent(ctx, promptMessage, additionalImageData)
	if err != nil {
		return "", err
	}
	return resp, nil
}
