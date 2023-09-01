package repository

import (
	"context"
	"database/sql"
)

type PostgreSQLClassicRepository struct {
	db *sql.DB
}

func NewPostgreSQLClassicRepository(db *sql.DB) *PostgreSQLClassicRepository {
	return &PostgreSQLClassicRepository{
		db: db,
	}
}

func (r *PostgreSQLClassicRepository) Migrate(ctx context.Context) error {
	query := `
    CREATE TABLE IF NOT EXISTS websites(
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL,
        rank INT NOT NULL
    );
    `

	_, err := r.db.ExecContext(ctx, query)
	return err
}