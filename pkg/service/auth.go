package service

import (
	"forwardcall/pkg/entity"
	"forwardcall/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) LoginUser(user entity.User) (int, error) {
	return s.repo.LoginUser(user)
}
