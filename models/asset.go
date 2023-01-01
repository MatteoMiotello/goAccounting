package models

type Asset struct {
	ID
	UserID        uint    `gorm:"not null"`
	Name          string  `json:"name" gorm:"not null"`
	InitialAmount float64 `json:"initialAmount"`
	User          User    `json:"user"`
	UpdatedCreated
}
