package models

type TransactionCategory struct {
	ID   uint
	Name string
	UpdatedCreated
	DeletedAt
}
