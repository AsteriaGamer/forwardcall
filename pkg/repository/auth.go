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

/*
	GetUser(username, password string) (entity.User, error)
	Функция обращается к БД и получает данные пользователя, по заданным логину и паролю.
	В ответ возвращается объект пользователя и ошибка.
*/
func (r *AuthRepository) GetUser(username string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password FROM users WHERE username=$1"
	err := r.db.Get(&user, query, username)

	return user, err
}
