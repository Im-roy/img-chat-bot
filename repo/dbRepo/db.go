package dbrepo

import "gorm.io/gorm"

type DbRepo struct {
	DB *gorm.DB
}
