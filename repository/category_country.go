package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-tests/models"
)

type CategoryCountryRepo interface {
	GetById(id int, country string) (*models.CategoryCountry, error)
}

type categoryCountryRepo struct{
	db *sql.DB
}
func NewCategoryCountryRepo(db *sql.DB) CategoryCountryRepo{
	return &categoryCountryRepo{}
}

func (cr categoryCountryRepo) 	GetById(id int, country string) (*models.CategoryCountry, error){
table := "category_country"
	query := fmt.Sprintf("SELECT * FROM %s WHERE category = %d AND country = %s", table,id, country)
 ctx := context.Background()
	row := cr.db.QueryRowContext(ctx, query)

	var category *models.CategoryCountry
	 if err := row.Scan(category.Id, category.Country, category.Category, category.Enabled); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("registro no existe  %s", err.Error())
        }
        return nil, err
    }

	return nil, nil

}