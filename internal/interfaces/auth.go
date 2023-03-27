package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
)

type Auth interface {
	Register(user models.User) (viewmodels.AuthedUser, error)
	GetCurrentUser(id string) (models.User, error)
}
