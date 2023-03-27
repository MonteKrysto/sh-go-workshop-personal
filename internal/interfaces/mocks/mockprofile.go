package mocks

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// Mock repository for Profile objects
type ProfileRepository struct {
	mock.Mock
}

func (m *ProfileRepository) Create(model models.Profile) (*viewmodels.Profile, error) {
	args := m.Called(model)
	return args.Get(0).(*viewmodels.Profile), args.Error(1)
}

func (m *ProfileRepository) GetAll() ([]viewmodels.Profile, error) {
	args := m.Called()
	return args.Get(0).([]viewmodels.Profile), args.Error(1)
}

func (m *ProfileRepository) GetById(id uuid.UUID) (*viewmodels.Profile, error) {
	args := m.Called(id)
	return args.Get(0).(*viewmodels.Profile), args.Error(1)
}

func (m *ProfileRepository) GetByUserId(id uuid.UUID) (*viewmodels.Profile, error) {
	args := m.Called(id)
	return args.Get(0).(*viewmodels.Profile), args.Error(1)
}

func (m *ProfileRepository) Update(model models.Profile, id uuid.UUID) (*viewmodels.Profile, error) {
	args := m.Called(model, id)
	return args.Get(0).(*viewmodels.Profile), args.Error(1)
}

func (m *ProfileRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
