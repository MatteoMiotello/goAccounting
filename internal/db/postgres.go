package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter struct {
}

func (p PostgresAdapter) GetAdapter() gorm.Dialector {
	dsn := fmt.Sprintf("postgres://%s@%s/%s?port=%s&password=%s",
		viper.Get("DB_USER"),
		viper.Get("DB_HOST"),
		viper.Get("DB_NAME"),
		viper.Get("DB_PORT"),
		viper.Get("DB_PASSWORD"))

	return postgres.Open(dsn)
}
