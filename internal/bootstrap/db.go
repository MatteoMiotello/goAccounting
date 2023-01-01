package bootstrap

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
)

type Db struct {
}

func (d Db) Init() error {
	err := db.Initialize(db.PostgresAdapter{})
	if err != nil {
		return err
	}
	return nil
}

func (d Db) InitMigration() error {
	err := db.Initialize(db.PostgresMigrationAdapter{})
	if err != nil {
		return err
	}
	return nil
}
