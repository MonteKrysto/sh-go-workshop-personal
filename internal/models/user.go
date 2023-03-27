package models

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	fmt.Println("BeforeCreate")
	uuid := uuid.New()

	scope.Statement.SetColumn("ID", uuid)

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	scope.Statement.SetColumn("Password", hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func (u *User) VerifyPassword(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}
