package repositories

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (cr CampaignRepository) Create(model models.Campaign) (*viewmodels.Campaign, error) {

	err := cr.db.Create(&model).Error
	var campaignVm viewmodels.Campaign

	campaignVm = viewmodels.Campaign{
		ID:          model.ID,
		UserId:      model.UserID,
		Name:        model.Name,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
	}

	return &campaignVm, err
}

func (cr CampaignRepository) GetAll() ([]viewmodels.Campaign, error) {
	var allCampaigns []viewmodels.Campaign

	// find all the surveys and put them in the allCampaigns slice
	err := cr.db.Find(&allCampaigns).Error
	return allCampaigns, err
}

func (cr CampaignRepository) GetById(id uuid.UUID) (*viewmodels.Campaign, error) {
	// using the models.Campaign{ID: id} filters out records in the query that have a timestamp in the deleted_at column
	campaign := &models.Campaign{
		BaseModelWithUUID: models.BaseModelWithUUID{
			ID: id,
		},
	}
	err := cr.db.First(&campaign).Error

	surveyVm := viewmodels.Campaign{
		ID:          campaign.ID,
		UserId:      campaign.UserID,
		Name:        campaign.Name,
		Description: campaign.Description,
		StartDate:   campaign.StartDate,
		EndDate:     campaign.EndDate,
	}
	return &surveyVm, err
}

func (cr CampaignRepository) Update(model models.Campaign, id uuid.UUID) (*viewmodels.Campaign, error) {
	var campaignVm viewmodels.Campaign

	err := cr.db.Clauses(clause.Returning{}).Model(&models.Campaign{
		BaseModelWithUUID: models.BaseModelWithUUID{ID: id}}).Updates(model).Error

	campaignVm = viewmodels.Campaign{
		ID:          id,
		Name:        model.Name,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
	}

	return &campaignVm, err
}

func (cr CampaignRepository) Delete(id uuid.UUID) error {
	campaign := &models.Campaign{
		BaseModelWithUUID: models.BaseModelWithUUID{
			ID: id,
		},
	}

	err := cr.db.Delete(&campaign).Error
	return err
}
