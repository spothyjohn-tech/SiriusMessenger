package database

import (
	"MesEdge/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Messages{}, &models.Group{}, &models.Channel{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
