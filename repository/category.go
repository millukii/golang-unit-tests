package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-tests/models"

)

type CategoryRepo interface {
	GetById(id int) (*models.Category, error)
}

type catRepo struct{
	db *sql.DB
}
func NewCategoryRepo(db *sql.DB) CategoryRepo{
	return &catRepo{}
}

func (c catRepo) 	GetById(id int) (*models.Category, error){

	query := "SELECT * FROM `categories` WHERE `id` = ?"
 ctx := context.Background()
	row := c.db.QueryRowContext(ctx, query)

	var category *models.Category
	 if err := row.Scan(category.Id, category.Name); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("registro no existe  %s", err.Error())
        }
        return nil, err
    }

	return nil, nil

}