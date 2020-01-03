package service

import (
	"cat-api/models"
	"cat-api/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type service struct {
	repo user.Repository
	secret []byte
}

type claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
func NewService(r user.Repository, secret []byte) user.Service {
	return &service{
		repo: r,
		secret: secret,
	}
}

func (s *service) Login(login models.UserLogin) (string, bool, error) {
	if login.UserName == "" || login.Password == "" {
		return "", false, fmt.Errorf("username or password are blank")
	}

	exists, err := s.repo.FindUserByCredentials(login)

	if err != nil {
		return "", false, err
	}

	if !exists {
		return "", false, fmt.Errorf("invalid username or password")
	}

	token, err := s.generateToken(login)

	if err != nil {
		return "", true, fmt.Errorf("error generating token")
	}

	return token, true, nil
}

func (s *service) generateToken(login models.UserLogin) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)

	claims := &claim{
		Username: login.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}