package repositories

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) Register(user models.User) (viewmodels.AuthedUser, error) {
	err := ar.db.Create(&user).Error
	var authedUser viewmodels.AuthedUser

	authedUser = viewmodels.AuthedUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	return authedUser, err
}

func (ar *AuthRepository) GetCurrentUser(id string) (models.User, error) {
	var user models.User
	err := ar.db.Where("id = ?", id).First(&user).Error

	return user, err
}
