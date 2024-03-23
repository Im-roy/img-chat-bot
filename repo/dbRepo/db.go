package dbrepo

import "gorm.io/gorm"

type DbRepo struct {
	db *gorm.DB
}
