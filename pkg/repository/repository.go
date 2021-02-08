package repository

import (
	"forwardcall/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	LoginUser(user entity.User) (int, error)
}

type ScheduleList interface {
}

type ContactList interface {
}

type Repository struct {
	Authorization
	ScheduleList
	ContactList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
