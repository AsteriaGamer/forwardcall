package service

import (
	"forwardcall/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) AuthUser(username, password string) (string, error) {

	// Получение данных пользователя из БД
	user, err := s.repo.GetUser(username)
	if err != nil {
		logrus.Errorf("Error, user %s doesn't find in db", user.Username)
		return "", err
	}

	// Проверка валидности пользовательского пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Errorf("Password user %s not match: %s", user.Username, err.Error())
		return "", err
	}

	// Если пользователь существует и пароль совпадает, генерируем jwt токен
	return user.Username, nil
}
