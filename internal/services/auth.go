package services

import (
	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
)

type AuthService struct {
	repository interfaces.Auth
}

// NewAuthService is a constructor that returns a new instance of AuthService
// It takes in a repository as a parameter
func NewAuthService(repository interfaces.Auth) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

// Register is responsible for registering a new user
// We return a view model of the authenticated user and/or an error
func (as AuthService) Register(user models.User) (viewmodels.AuthedUser, error) {
	return as.repository.Register(user)
}

func (as AuthService) GetCurrentUser(id string) (models.User, error) {
	return as.repository.GetCurrentUser(id)
}
