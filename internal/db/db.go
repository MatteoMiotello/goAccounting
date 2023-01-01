package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(adapter dbAdapter) error {
	db, err := gorm.Open(adapter.GetAdapter(), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
