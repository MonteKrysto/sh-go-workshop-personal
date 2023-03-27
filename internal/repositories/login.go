package repositories

import (
	"fmt"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{
		db: db,
	}
}

func (lr *LoginRepository) ValidateLogin(login models.Login) (string, string, error) {
	var user models.User

	err := lr.db.Where("username = ?", login.Username).First(&user).Error
	if err != nil {
		return "", "", err
	}
	err = user.VerifyPassword(login.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", err
	}

	return fmt.Sprintf("%v", user.ID), fmt.Sprintf("%v", user.Password), nil
}
