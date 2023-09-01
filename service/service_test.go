package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-tests/models"
	"go-tests/service"
	"go-tests/mocks"
)

		// Crear un mock del servicio
func TestMockService_GetById(t *testing.T) {


		t.Run("Prueba happy path", func (t *testing.T){
			// Crear mocks de los repositorios
			categoryRepo := new(mocks.CategoryRepo)
			categoryCountryRepo := new(mocks.CategoryCountryRepo)

			// Crear un mock del servicio con los repositorios mockeados
			mockService := service.NewService(categoryRepo, categoryCountryRepo)
					// Configurar comportamiento de los mocks
			categoryRepo.On("GetById", 1).Return(&models.Category{Id: 1, 	Name: "esports"}, nil)
			categoryCountryRepo.On("GetById", 1, "US").Return(&models.CategoryCountry{Id: 1, Enabled: true}, nil)

			// Caso de prueba: Categoría habilitada
			category, err := mockService.GetById(1, "US")
			assert.NoError(t, err)
			assert.NotNil(t, category)
			assert.Equal(t, 1, category.Id)

			// Asegurarse de que todas las expectativas se cumplan
			categoryRepo.AssertExpectations(t)
			categoryCountryRepo.AssertExpectations(t)
		})


}
 