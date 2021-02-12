package service

import (
	"forwardcall/pkg/repository"
)

type Authorization interface {
	AuthUser(username, password string) (string, error)
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
