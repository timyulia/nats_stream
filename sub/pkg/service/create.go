package service

import (
	"github.com/go-playground/validator/v10"
	"nats"
	cache "nats/pkg/cashe"
	repository "nats/pkg/repository"
)

type CreatorService struct {
	repo  repository.Creator
	inMem cache.InMemory
}

func NewCreatorService(repo repository.Creator, inMem cache.InMemory) *CreatorService {
	return &CreatorService{repo: repo, inMem: inMem}
}

func (s *CreatorService) Create(ord nats.Order) error {
	err := s.validate(ord)
	if err != nil {
		return err
	}
	err = s.repo.Create(&ord)
	if err != nil {
		return err
	}
	err = s.inMem.SaveOrder(&ord)
	return err
}

func (s *CreatorService) validate(ord nats.Order) error {
	validate := validator.New()
	return validate.Struct(ord)
}
