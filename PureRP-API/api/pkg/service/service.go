package service

import (
	"api"
	"api/pkg/repository"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
}

type Character interface {
}

type Service struct {
	Authorization
	Character
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
