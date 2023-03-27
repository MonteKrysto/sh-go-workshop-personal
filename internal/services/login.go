package services

import (
	"encoding/json"
	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	token "github.com/SpringCare/sh-go-workshop/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repository interfaces.Login
}

// NewLoginService is a constructor that returns a new instance of LoginService
// It takes in a repository as a parameter
func NewLoginService(repository interfaces.Login) *LoginService {
	return &LoginService{
		repository: repository,
	}
}

// ValidateLogin is responsible for logging in a user
// We return a view model of the authenticated user and/or an error
func (ls LoginService) ValidateLogin(login models.Login) (string, string, error) {
	var user models.User

	userId, pwd, err := ls.repository.ValidateLogin(login)
	if err != nil {
		return "", "", err
	}
	json.Unmarshal([]byte(userId), &user)

	err = VerifyPassword(login.Password, pwd) //user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", err
	}

	token, err := token.GenerateToken(userId)

	if err != nil {
		return "", "", err
	}

	return token, "", nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
