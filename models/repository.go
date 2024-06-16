package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h24s_17-backend/utils"
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

// なかったら作る
func (r *Repository) ReadLatestToken() (*Token, error) {
	var token Token
	query := "SELECT token FROM tokens WHERE DATE_ADD(created_at, INTERVAL 5 MINUTE) <= NOW()"
	err := r.db.Get(&token, query)
	if err == nil {
		return &token, nil
	}

	newToken, err := utils.GenerateRandomToken(12)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec("INSERT INTO tokens (token) VALUES (?)", newToken)
	if err != nil {
		return nil, err
	}
	return &Token{Token: newToken}, nil
}
