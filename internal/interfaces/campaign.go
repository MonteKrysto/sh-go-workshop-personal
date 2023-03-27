package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
)

type Campaign interface {
	Create(campaign models.Campaign) (viewmodels.Campaign, error)
	//GetById(id uuid.UUID) (viewmodels.Campaign, error)
	//GetByUserId(id uuid.UUID) (viewmodels.Campaign, error)
	//Update(campaign models.Campaign, id uuid.UUID) error
	//Delete(id uuid.UUID) error
}
