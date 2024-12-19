package model

import (
	"mime/multipart"
	"time"
)

type PromptImageModel struct {
	ExtensionName string
	Data          []byte
}

type RequestModel struct {
	Prompt string `json:"data"`
}

type UserFilePathMapping struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int
	FilePath  string    `gorm:"not null;uniqueIndex:idx_user_id_filepath"`
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type FileDetailsModel struct {
	Header *multipart.FileHeader
	Data   multipart.File
}
