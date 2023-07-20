package repository

import (
	"api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
}

type Character interface {
}

type Repository struct {
	Authorization
	Character
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
