package repositories

import (
	"fmt"

	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockery --name=ProfileRepository --output=./mocks --case=underscore
type ProfileRepository struct {
	db *gorm.DB
}

//go:generate mockery --name=NewProfileRepository --output=./mocks --case=underscore
func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (pr ProfileRepository) GetAll() ([]viewmodels.Profile, error) {
	var profiles []models.Profile
	err := pr.db.Find(&profiles).Error

	var profilesVM []viewmodels.Profile
	for _, profile := range profiles {
		//profileVM := viewmodels.Profile{
		//	ID:      profile.ID,
		//	UserId:  profile.UserId,
		//	Street:  profile.Street,
		//	City:    profile.City,
		//	State:   profile.State,
		//	ZipCode: profile.ZipCode,
		//}

		profileVM := pr.convertToViewModel(profile)
		profilesVM = append(profilesVM, *profileVM)
	}

	return profilesVM, err
}

func (pr ProfileRepository) GetById(id uuid.UUID) (*viewmodels.Profile, error) {
	profile := models.Profile{}
	err := pr.db.First(&profile, id).Error

	return pr.convertToViewModel(profile), err
}

func (pr ProfileRepository) GetByUserId(id uuid.UUID) (*viewmodels.Profile, error) {
	profile := models.Profile{}
	err := pr.db.Find(&profile, "user_id = ?", id).Error

	return pr.convertToViewModel(profile), err
}

func (pr ProfileRepository) Create(model models.Profile) (*viewmodels.Profile, error) {
	fmt.Println("asdf")
	err := pr.db.FirstOrCreate(&model).Error

	return pr.convertToViewModel(model), err
}

func (pr ProfileRepository) Update(model models.Profile, id uuid.UUID) (*viewmodels.Profile, error) {
	//var profileVM viewmodels.Profile

	err := pr.db.Clauses(clause.Returning{}).Model(&models.Profile{
		BaseModelWithUUID: models.BaseModelWithUUID{ID: id}}).Updates(model).Error
	//if err != nil {
	//	// We could skip this and just return the model that was supplied instead of an empty model and the error
	//	return &profileVM, err
	//}

	// If there is an error we'll return the model that was supplied as well as the error
	return pr.convertToViewModel(model), err
}

func (pr ProfileRepository) Delete(id uuid.UUID) error {
	err := pr.db.Delete(&models.Profile{
		BaseModelWithUUID: models.BaseModelWithUUID{
			ID: id,
		}}).Error

	return err
}

func (pr ProfileRepository) convertToViewModel(model models.Profile) *viewmodels.Profile {
	return &viewmodels.Profile{
		ID:      model.ID,
		UserId:  model.UserId,
		Street:  model.Street,
		City:    model.City,
		State:   model.State,
		ZipCode: model.ZipCode,
	}
}
