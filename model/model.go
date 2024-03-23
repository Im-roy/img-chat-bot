package model

import "time"

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
	FilePath  string    `gorm:"size:255;not null"`
	IsActive  bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
