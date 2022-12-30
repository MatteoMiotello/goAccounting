package db

import (
	"github.com/MatteoMiotello/goAccounting/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(adapter dbAdapter) error {
	db, err := gorm.Open(adapter.GetAdapter(), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Asset{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
