package migrations

import (
	"database/sql"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(up1, down1)
}

func up1(tx *sql.Tx) error {
	if !db.DB.Migrator().HasTable(&models.User{}) {
		err := db.DB.Migrator().CreateTable(&models.User{})
		if err != nil {
			return err
		}
	}

	if !db.DB.Migrator().HasTable(&models.Asset{}) {
		err := db.DB.Migrator().CreateTable(&models.Asset{})
		if err != nil {
			return err
		}

	}

	return nil
}

func down1(tx *sql.Tx) error {
	err := db.DB.Migrator().DropTable(&models.Asset{})
	if err != nil {
		return err

	}

	err = db.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}
	return nil
}
