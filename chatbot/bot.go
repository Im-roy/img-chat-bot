package chatbot

import (
	"context"
	aimodel "img-chat-bot/AIModel"
	datamapper "img-chat-bot/dataMapper"
	"img-chat-bot/model"
	dbrepo "img-chat-bot/repo/dbRepo"
	fileRepo "img-chat-bot/repo/fileRepo"
	"log"
	"os"
)

type ChatBot struct {
	AIModel  aimodel.AiModel
	FileRepo fileRepo.FileRepo
	DbRepo   dbrepo.DbRepo
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

func (cb *ChatBot) SaveUserImage(ctx context.Context, file model.FileDetailsModel, userID int) error {

	uploadDir := "./images/"
	err := cb.FileRepo.SetDirectory(uploadDir)
	if err != nil {
		return err
	}

	err = cb.FileRepo.SaveFile(ctx, file)
	if err != nil {
		return err
	}

	userFileMappings := datamapper.GetUserFileMappingsGormModel(file, userID)
	err = cb.DbRepo.CreateMappings(ctx, userFileMappings)
	if err != nil {
		return err
	}
	return nil
}
