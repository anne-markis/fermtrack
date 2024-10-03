package app

import (
	"errors"
	"fmt"

	"github.com/anne-markis/fermtrack/internal/app/domain"
	"github.com/anne-markis/fermtrack/internal/utils"
)

// TODO test
type AuthService struct {
	userRepo domain.UserRepository
}

type Authenicator interface {
	Login(username, password string) (string, error)
}

func NewAuthService(userRepo domain.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (a *AuthService) Login(username, password string) (string, error) {
	user, err := a.userRepo.FindByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate jwt token: %s", err.Error()))
	}

	return token, nil
}
