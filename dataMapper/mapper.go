package datamapper

import (
	"img-chat-bot/model"
	"time"
)

func GetUserFileMappingsGormModel(file model.FileDetailsModel, userID int) model.UserFilePathMapping {
	return model.UserFilePathMapping{
		UserID:    userID,
		FilePath:  file.Header.Filename,
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}
