package chatbot

import (
	"context"
	aimodel "img-chat-bot/AIModel"
	datamapper "img-chat-bot/dataMapper"
	"img-chat-bot/model"
	dbrepo "img-chat-bot/repo/dbRepo"
	fileRepo "img-chat-bot/repo/fileRepo"
	"img-chat-bot/utils"
	"os"
)

type ChatBot struct {
	AIModel  aimodel.AiModel
	FileRepo fileRepo.FileRepo
	DbRepo   dbrepo.DbRepo
}

func (cb *ChatBot) GenerateResponse(ctx context.Context, promptMessage string, userID int) (string, error) {

	userImageFilePathMappings, err := cb.GetUserImages(ctx, userID)
	if err != nil {
		return "", err
	}

	additionalImageData := []model.PromptImageModel{}
	for _, mappings := range userImageFilePathMappings {
		pathToImage := "images/" + mappings.FilePath
		imgData, err := os.ReadFile(pathToImage)
		if err != nil {
			return "", err
		}
		additionalImageData = append(additionalImageData, model.PromptImageModel{
			ExtensionName: utils.ExtractExtension(mappings.FilePath),
			Data:          imgData,
		})
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

func (cb *ChatBot) GetUserImages(ctx context.Context, userID int) ([]model.UserFilePathMapping, error) {

	userFilepathMappings, err := cb.DbRepo.GetMappings(ctx, userID)
	if err != nil {
		return []model.UserFilePathMapping{}, err
	}
	return userFilepathMappings, nil
}
