package repository

import (
	"Go-rest-api/internal/domain"
	"context"

	"gorm.io/gorm"
)

type UserDBRepo struct {
	db *gorm.DB
}

func NewUserDBRepo(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{
		db: db,
	}
}

func (repo UserDBRepo) GetByID(ctx context.Context, ID int) (domain.User, error) {
	var user domain.User
	err := repo.db.Where("id = ?", ID).Take(&user).Error
	return user, err
}