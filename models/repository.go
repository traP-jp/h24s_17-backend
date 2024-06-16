package models

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

type Token struct {
	Token string `json:"token,omitempty" db:"token"`
}

func ConnectRepository() (*Repository, error) {
	db, err := Connect()

	return &Repository{db: db}, err
}

func (r *Repository) Migrate() error {
	return Migrate(r.db)
}

func (r *Repository) CreateToken(token string) (Token, error) {
	query := "INSERT INTO tokens (token) VALUES (\"" + token + "\");"

	_, err := r.db.Exec(query)

	return Token{token}, err
}

func (r *Repository) ReadTokens() ([]Token, error) {
	var tokens []Token
	query := "SELECT token FROM tokens"
	err := r.db.Select(&tokens, query)

	return tokens, err
}
