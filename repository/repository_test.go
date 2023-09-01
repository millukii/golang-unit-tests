package repository_test

/* import (
	"database/sql"
	"go-tests/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

)

func TestRepository(t *testing.T) {

	t.Run("Test Case ok", func(t *testing.T) {

db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM country_categories").WillReturnError(sql.ErrNoRows)

	//mock.ExpectExec("SELECT * FROM country_categories").WillReturnResult(sqlmock.NewResult(1, 1))
	//WillReturnError(sql.ErrNoRows)
	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	//crear instancia del repo

	repo :=repository.NewCategoryRepo(db)
	// now we execute our method
		cats ,err := repo.GetById( 2, "chile")
		if  err != nil {
			t.Errorf("error was not expected while updating stats: %s", err)
		}
		assert.Nil(t,cats)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	})
} */