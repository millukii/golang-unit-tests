package mocks

import (
	"go-tests/models"

	"github.com/stretchr/testify/mock"

)

// Service es una estructura mock para la interfaz Service.
type Service struct {
	mock.Mock
}

// GetById provee una implementación mock de GetById en la interfaz Service.
func (m *Service) GetById(id int, country string) (*models.Category, error) {
	args := m.Called(id, country)
	return args.Get(0).(*models.Category), args.Error(1)
}