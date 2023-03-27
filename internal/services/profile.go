package services

import (
	"fmt"

	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
)

type ProfileService struct {
	repository interfaces.BaseWithGetByUserId[models.Profile, viewmodels.Profile]
}

func NewProfileService(repository interfaces.BaseWithGetByUserId[models.Profile, viewmodels.Profile]) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}

func (ps ProfileService) Create(model models.Profile) (*viewmodels.Profile, error) {
	fmt.Println("here")
	return ps.repository.Create(model)
}

func (ps ProfileService) GetAll() ([]viewmodels.Profile, error) {
	return ps.repository.GetAll()
}

func (ps ProfileService) GetById(id uuid.UUID) (*viewmodels.Profile, error) {
	return ps.repository.GetById(id)
}

func (ps ProfileService) GetByUserId(id uuid.UUID) (*viewmodels.Profile, error) {
	return ps.repository.GetByUserId(id)
}

func (ps ProfileService) Update(model models.Profile, id uuid.UUID) (*viewmodels.Profile, error) {
	//return ps.repository.Update(model, id)
	p, e := ps.repository.Update(model, id)
	fmt.Println("p: ", p)
	return p, e
}

func (ps ProfileService) Delete(id uuid.UUID) error {
	return ps.repository.Delete(id)
}
