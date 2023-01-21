package repo

import (
	"context"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepo BaseRepo

func NewUserRepo() *UserRepo {
	return new(UserRepo)
}

func (u UserRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, db.DB)
}

func (u UserRepo) GetById(ctx context.Context, id int64) (*models.User, error) {
	return models.FindUser(ctx, db.DB, id)
}

func (u UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return models.Users(qm.Where("email=?", email)).One(ctx, db.DB)
}

func (u UserRepo) Insert(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, db.DB, boil.Infer())
}
