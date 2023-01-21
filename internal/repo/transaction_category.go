package repo

import (
	"context"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TransactionCategoryRepo BaseRepo

func NewTransactionCategoryRepo() *TransactionCategoryRepo {
	return new(TransactionCategoryRepo)
}

func (t TransactionCategoryRepo) GetActiveByName(ctx context.Context, name string) (*models.TransactionCategory, error) {
	return models.TransactionCategories(qm.Where("name = ? date_deleted is null", name)).One(ctx, db.DB)
}
