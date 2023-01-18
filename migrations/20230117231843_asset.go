package migrations

import (
	"database/sql"
	"github.com/MatteoMiotello/go-sqlbuilder"
	"github.com/MatteoMiotello/go-sqlbuilder/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAsset, downAsset)
}

func upAsset(tx *sql.Tx) error {
	assetQuery := sqlbuilder.CreateTable("public.assets").
		PKColumn().
		FKColumn("public.users", "user_id", false).
		Column("name", types.Varchar+"(50)", false).
		Column("description", types.Varchar+"(255)", true).
		Column("initial_amount", types.Decimal+"(2,16)", false).
		DeletedColumn().
		UpdatedColumn().
		CreatedColumn().
		String()

	_, err := tx.Exec(assetQuery)
	if err != nil {
		return err
	}

	return nil
}

func downAsset(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE public.assets")
	if err != nil {
		return err
	}
	return nil
}
