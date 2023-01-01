package models

type Transaction struct {
	ID
	AssetID               uint
	TransactionCategoryID uint
	Amount                float64
	Description           string
	TransactionCategory   TransactionCategory
	Asset                 Asset
	UpdatedCreated
	DeletedAt
}
