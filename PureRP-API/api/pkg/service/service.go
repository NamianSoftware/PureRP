package service

import "api/pkg/repository"

type Authorization interface {
}

type Character interface {
}

type Service struct {
	Authorization
	Character
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
