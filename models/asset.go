package models

type Asset struct {
	PrimaryKey
	UserID        uint    `json:"userId" gorm:"not null"`
	Name          string  `json:"name" gorm:"not null"`
	InitialAmount float64 `json:"initialAmount"`
	User          User    `json:"user,omitempty"`
	UpdatedCreated
}
