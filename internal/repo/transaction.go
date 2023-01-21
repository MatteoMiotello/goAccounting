package repo

import (
	"context"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TransactionRepo BaseRepo

func NewTransactionRepo() *TransactionRepo {
	return new(TransactionRepo)
}

func (t TransactionRepo) Insert(ctx context.Context, model *models.Transaction) error {
	return model.Insert(ctx, db.DB, boil.Infer())
}
