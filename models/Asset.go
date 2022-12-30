package models

import (
	"time"
)

type Asset struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	InitialAmount float64   `json:"initial_amount"`
	CreatedAt     time.Time `json:"created_at,string,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,string,omitempty"`
}
