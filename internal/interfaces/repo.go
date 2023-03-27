package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/google/uuid"
)

//go:generate mockery --name=Model --output=./mocks --case=underscore
type Model interface {
	models.Campaign | models.Profile
}

//go:generate mockery --name=ViewModel --output=./mocks --case=underscore
type ViewModel interface {
	viewmodels.Campaign | viewmodels.Profile
}

//go:generate mockery --name=BaseRepository --output=./mocks --case=underscore
type BaseRepository[M Model, V ViewModel] interface {
	Create(model M) (*V, error)
	GetAll() ([]V, error)
	GetById(id uuid.UUID) (*V, error)
	Update(model M, id uuid.UUID) (*V, error)
	Delete(id uuid.UUID) error
}

//go:generate mockery --name=BaseWithGetByUserId --output=./mocks --case=underscore
type BaseWithGetByUserId[M Model, V ViewModel] interface {
	BaseRepository[M, V]
	GetByUserId(id uuid.UUID) (*V, error)
}

//type WithDelete[M Model, V ViewModel] interface {
//	BaseRepository[M, V]
//	Delete(id uuid.UUID) error
//}
