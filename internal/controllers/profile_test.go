package controllers

import (
	"bytes"
	"github.com/stretchr/testify/assert"

	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SpringCare/sh-go-workshop/internal/interfaces/mocks"
)

func TestCreateProfile(t *testing.T) {
	// Arrange
	mockProfileService := new(services.ProfileService)
	mockProfileRepository := new(mocks.ProfileRepository)
	mockProfileController := NewProfileController(*mockProfileService)

	profile := models.Profile{
		BaseModelWithUUID: models.BaseModelWithUUID{
			ID: uuid.New(),
		},
		UserId:  uuid.New(),
		Street:  "Test Street",
		City:    "test city",
		State:   "SC",
		ZipCode: "12345",
	}

	// Set up expected behavior of mock objects
	mockProfileRepository.On("Create", profile).Return(&viewmodels.Profile{}, nil)

	// Create a Gin HTTP request
	body := []byte(`{"id":"` + profile.ID.String() + `","user_id":"` + profile.UserId.String() + `","street":"` + profile.Street + `","city":"` + profile.City + `","state":"` + profile.State + `","zip_code":"` + profile.ZipCode + `"}`)
	req, _ := http.NewRequest("POST", "/profile", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	mockProfileController.Create(c)

	// Assert
	mockProfileRepository.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}
