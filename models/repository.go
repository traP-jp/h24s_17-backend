package models

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func ConnectRepository() (*Repository, error) {
	db, err := Connect()

	return &Repository{db: db}, err
}

func (r *Repository) Migrate() error {
	return Migrate(r.db)
}
