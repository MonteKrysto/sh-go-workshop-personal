package services

import (
	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
)

type CampaignService struct {
	repository interfaces.BaseRepository[models.Campaign, viewmodels.Campaign]
}

func NewCampaignService(repository interfaces.BaseRepository[models.Campaign, viewmodels.Campaign]) *CampaignService {
	return &CampaignService{
		repository: repository,
	}
}

func (cs *CampaignService) Create(model models.Campaign) (*viewmodels.Campaign, error) {
	return cs.repository.Create(model)
}

func (cs *CampaignService) GetAll() ([]viewmodels.Campaign, error) {
	return cs.repository.GetAll()
}

func (cs *CampaignService) GetById(id uuid.UUID) (*viewmodels.Campaign, error) {
	return cs.repository.GetById(id)
}

func (cs *CampaignService) Update(campaign models.Campaign, id uuid.UUID) (*viewmodels.Campaign, error) {
	return cs.repository.Update(campaign, id)
}

func (cs *CampaignService) Delete(id uuid.UUID) error {
	return cs.repository.Delete(id)
}
