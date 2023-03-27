package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
)

type Login interface {
	ValidateLogin(login models.Login) (string, string, error)
}
