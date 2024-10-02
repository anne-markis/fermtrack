package app

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/anne-markis/fermtrack/internal/app/domain"
	"github.com/anne-markis/fermtrack/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

// TODO test
type AuthService struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

type Authenicator interface {
	Login(username, password string) (string, error)
}

func NewAuthService(userRepo domain.UserRepository) *AuthService {
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	return &AuthService{userRepo: userRepo, jwtSecret: jwtKey}
}

func (a *AuthService) Login(username, password string) (string, error) {
	user, err := a.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := a.generateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) generateJWT(username string) (string, error) {
	if a.jwtSecret == "" {
		return "", fmt.Errorf("jwt not set up")
	}
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.jwtSecret)
}
