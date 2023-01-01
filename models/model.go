package models

import (
	"gorm.io/gorm"
	"time"
)

type UpdatedCreated struct {
	CreatedAt time.Time `json:"createdAt,string,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,string,omitempty"`
}

type PrimaryKey struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,string,omitempty"`
}
