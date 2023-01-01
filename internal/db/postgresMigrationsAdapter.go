package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresMigrationAdapter struct {
}

func (p PostgresMigrationAdapter) GetAdapter() gorm.Dialector {
	dsn := fmt.Sprintf("postgres://%s@%s/%s?port=%s&password=%s",
		viper.Get("MIGRATION_DB_USER"),
		viper.Get("MIGRATION_DB_HOST"),
		viper.Get("MIGRATION_DB_NAME"),
		viper.Get("MIGRATION_DB_PORT"),
		viper.Get("MIGRATION_DB_PASSWORD"))

	return postgres.Open(dsn)
}
