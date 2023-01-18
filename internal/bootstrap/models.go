package bootstrap

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InitModels() {
	boil.SetDB(db.DB)
}
