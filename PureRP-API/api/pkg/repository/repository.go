package repository

type Authorization interface {
}

type Character interface {
}

type Repository struct {
	Authorization
	Character
}

func NewRepository() *Repository {
	return &Repository{}
}
