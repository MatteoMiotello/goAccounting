package db

import "gorm.io/gorm"

type dbAdapter interface {
	GetAdapter() gorm.Dialector
}
