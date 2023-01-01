package migrations

import (
	"database/sql"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upTransaction, downTransaction)
}

func upTransaction(tx *sql.Tx) error {
	if !db.DB.Migrator().HasTable(&models.TransactionCategory{}) {
		err := db.DB.Migrator().CreateTable(&models.TransactionCategory{})
		if err != nil {
			return err
		}
	}

	if !db.DB.Migrator().HasTable(&models.Transaction{}) {
		err := db.DB.Migrator().CreateTable(&models.Transaction{})
		if err != nil {
			return err
		}
	}
	return nil
}

func downTransaction(tx *sql.Tx) error {
	db.DB.Migrator().DropTable(&models.Transaction{})
	db.DB.Migrator().DropTable(&models.TransactionCategory{})
	return nil
}
