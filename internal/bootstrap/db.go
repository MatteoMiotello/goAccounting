package bootstrap

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
)

func InitDb() {
	err := db.Initialize(db.PostgresAdapter{})
	if err != nil {
		panic("error initializing db: " + err.Error())
	}
}

func InitDbMigration() {
	err := db.Initialize(db.PostgresMigrationAdapter{})
	if err != nil {
		panic("error initializing db: " + err.Error())
	}
}
