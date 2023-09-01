package service

import (
	"errors"
	"fmt"
	"go-tests/models"
	"go-tests/repository"
)


type Service interface {
	GetById(id int, country string) (*models.Category, error)
}

type svc struct {
	//dependencias
	repoCategory repository.CategoryRepo
	repoCategoryCountry repository.CategoryCountryRepo
}

func NewService(repoCategory repository.CategoryRepo, repoCategoryCountry repository.CategoryCountryRepo) Service {

	return &svc{
		repoCategory: repoCategory,
		repoCategoryCountry: repoCategoryCountry,
	}
}

func (s svc) GetById(id int, country string) (*models.Category, error) {

	//acceso a los datos
	//obtener categoria y sus atributos
	cat, err := s.repoCategory.GetById(id)

	if err != nil{
		return nil, err
	}
	//obtener categoria por pais
	categoryCountry, err := s.repoCategoryCountry.GetById(id, country)
	//logica de negocio -- reglas
	//validar si la categoria está habilitada
if err !=nil{
	return nil, err
}

if !categoryCountry.Enabled {
	return nil, fmt.Errorf("%s", errors.New("category is disable for this country"))
}

	return cat, nil
}