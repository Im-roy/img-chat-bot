package dbrepo

import (
	"context"
	"img-chat-bot/model"

	"gorm.io/gorm/clause"
)

func (db *DbRepo) CreateMappings(ctx context.Context, mappings model.UserFilePathMapping) error {
	err := db.DB.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "file_path"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_active"}),
	}).Create(&mappings).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DbRepo) GetMappings(ctx context.Context, userID int) ([]model.UserFilePathMapping, error) {

	userFileMappings := []model.UserFilePathMapping{}
	err := db.DB.Debug().Find(&userFileMappings).Where("user_id = ? and is_active = true", userID).Error
	if err != nil {
		return []model.UserFilePathMapping{}, err
	}

	return userFileMappings, nil
}
