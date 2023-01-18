package migrations

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upFirstMigration, downFirstMigration)
}

func upFirstMigration(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.users").
		PKColumn().
		Column("username", types.Varchar, false).
		Column("email", types.Varchar, false).
		Column("password", types.Varchar, false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downFirstMigration(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
