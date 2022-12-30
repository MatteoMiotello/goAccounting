package models

import (
	"gorm.io/gorm"
	"time"
)

type model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt,string,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,string,omitempty"`
}

type deletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,string,omitempty"`
}
