package models

type Transaction struct {
	PrimaryKey
	AssetID               uint
	TransactionCategoryID uint
	Amount                float64
	Description           string
	TransactionCategory   TransactionCategory
	Asset                 Asset
	UpdatedCreated
	DeletedAt
}
