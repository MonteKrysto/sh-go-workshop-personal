package services

import (
	"github.com/SpringCare/sh-go-workshop/internal/interfaces/mocks"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProfileService_Create(t *testing.T) {
	repo := &mocks.BaseWithGetByUserId[models.Profile, viewmodels.Profile]{} //  repomocks.ProductRepositoryInterface{}
	vm := viewmodels.Profile{
		UserId:  uuid.New(),
		Street:  "123 Main St",
		City:    "San Francisco",
		State:   "CA",
		ZipCode: "94105",
	}
	repo.On("Create", mock.AnythingOfType("models.Profile")).
		Return(&vm, nil).
		Once()

	service := NewProfileService(repo) // services.NewProductService(repo)

	err, _ := service.Create(models.Profile{
		UserId:  uuid.New(),
		Street:  "123 Main St",
		City:    "San Francisco",
		State:   "CA",
		ZipCode: "94105",
	})

	//err := service.Insert("2f1afe98-63c4-4f59-bcaf-1df835602bdb", models.InsertProductDTO{
	//	Name:  "Macbook",
	//	Price: 20500,
	//	Stock: 10,
	//})

	assert.Nil(t, err)

}
