package repository

import (
	"api"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(user api.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (first_name, last_name, email, password_hash) values ($1, $2, $3, $4) RETURNING id;",
		usersTable,
	)
	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
