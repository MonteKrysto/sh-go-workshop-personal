package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
)

type Profile interface {
	Create(profile models.Profile) (viewmodels.Profile, error)
	GetById(id uuid.UUID) (viewmodels.Profile, error)
	GetByUserId(id uuid.UUID) (viewmodels.Profile, error)
	Update(profile models.Profile, id uuid.UUID) error
	Delete(id uuid.UUID) error
}
