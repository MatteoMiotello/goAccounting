package migrations

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upTransactions, downTransactions)
}

func upTransactions(tx *sql.Tx) error {
	query := sqlbuilder.CreateTable("public.transaction_categories").
		PKColumn().
		Column("name", types.Varchar+"(50)", false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	query = sqlbuilder.CreateTable("public.transactions").
		PKColumn().
		FKColumn("public.assets", "asset_id", false).
		FKColumn("public.transaction_categories", "transaction_category_id", false).
		Column("amount", types.Decimal+"(2,16)", false).
		Column("description", types.Varchar+"(255)", true).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func downTransactions(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.transactions")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE public.transaction_categories")
	if err != nil {
		return err
	}
	return nil
}
