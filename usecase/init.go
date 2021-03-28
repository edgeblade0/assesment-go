package usecase

import (
	"assesment/repository"
)

type uc struct {
	repo repository.Repository
}

type Usecase interface {
	GetTracking() (string, error)
}

func InitUsecase(r repository.Repository) Usecase {
	return &uc{
		repo: r,
	}
}
