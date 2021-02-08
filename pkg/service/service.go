package service

import (
	"forwardcall/pkg/entity"
	"forwardcall/pkg/repository"
)

type Authorization interface {
	LoginUser(user entity.User) (int, error)
}

type ScheduleList interface {
}

type ContactList interface {
}

type Service struct {
	Authorization
	ScheduleList
	ContactList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
