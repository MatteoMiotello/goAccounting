package models

type Asset struct {
	model
	Name          string  `json:"name"`
	InitialAmount float64 `json:"initialAmount"`
}
