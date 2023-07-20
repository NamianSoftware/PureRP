package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Character interface {
}

type Repository struct {
	Authorization
	Character
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
