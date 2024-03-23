package chatbot

import (
	"context"
	aimodel "img-chat-bot/AIModel"
	"img-chat-bot/model"
	"log"
	"os"
)

type ChatBot struct {
	AIModel aimodel.AiModel
}

func (cb *ChatBot) GenerateResponse(ctx context.Context, promptMessage string, userID int) (string, error) {

	pathToImage1 := "images/abhi.jpeg"
	imgData1, err := os.ReadFile(pathToImage1)
	if err != nil {
		log.Fatal(err)
	}

	pathToImage2 := "images/announcement.jpg"
	imgData2, err := os.ReadFile(pathToImage2)
	if err != nil {
		log.Fatal(err)
	}
	additionalImageData := []model.PromptImageModel{
		{
			ExtensionName: "jpeg",
			Data:          imgData1,
		},
		{
			ExtensionName: "jpg",
			Data:          imgData2,
		},
	}
	aiResponse, err := cb.AIModel.GenerateResponse(ctx, promptMessage, additionalImageData)
	return aiResponse, nil
}
