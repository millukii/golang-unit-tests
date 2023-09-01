package service_test

import (
	"go-tests/models"
	"go-tests/service"
	"go-tests/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestServiceMockManual(t *testing.T) {

	var repoCategoryMock RepoCategoryMock
	var repoCategoryCountryMock RepoCategoryCountryMock
	id := 1
	country := "chile"
	//dependencias mocks de los repositorios
	t.Run("happy path", func(t *testing.T) {

		//instancia al servicio
		svc := service.NewService(repoCategoryMock, repoCategoryCountryMock)

	 cat , err :=	svc.GetById(id, country)

	 //no deben existir errores
	 assert.Nil(t,err)
	 //categoria no debe ser nula
	 assert.NotNil(t, cat)
	 //la categoria deberia estar habilitada para el pais


	})
}
//Mocks
 type RepoCategoryMock struct{

 }
 func(rcMock RepoCategoryMock)	GetById(id int) (*models.Category, error){
	return &models.Category{
		Id: 1,
		Name: "ebooks",
	},nil
 }
 type RepoCategoryCountryMock struct{

 }
  func(rcMock RepoCategoryCountryMock)	GetById(id int, country string) (*models.CountriesCategory, error){
	return &models.CountriesCategory{
		Id: 1,
		Country: "chile",
		Category: 1,
		Enabled: true,
	},nil
 }

		// Crear un mock del servicio
func TestMockService_GetById(t *testing.T) {
	// Crear mocks de los repositorios
	categoryRepo := new(mocks.CategoryRepo)
	categoryCountryRepo := new(mocks.CategoryCountryRepo)

	// Crear un mock del servicio con los repositorios mockeados
	mockService := NewService(categoryRepo, categoryCountryRepo)

	// Configurar comportamiento de los mocks
	categoryRepo.On("GetById", 1).Return(&Category{ID: 1}, nil)
	categoryCountryRepo.On("GetById", 1, "US").Return(&Category{ID: 1, Enabled: true}, nil)
	categoryCountryRepo.On("GetById", 2, "US").Return(&Category{ID: 2, Enabled: false}, nil)

	// Caso de prueba: Categoría habilitada
	category, err := mockService.GetById(1, "US")
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, 1, category.ID)

	// Caso de prueba: Categoría deshabilitada
	category, err = mockService.GetById(2, "US")
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Contains(t, err.Error(), "disabled")

	// Asegurarse de que todas las expectativas se cumplan
	categoryRepo.AssertExpectations(t)
	categoryCountryRepo.AssertExpectations(t)
}
 