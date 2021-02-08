package repository

import (
	"forwardcall/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) LoginUser(user entity.User) (int, error) {
	return 0, nil
}
