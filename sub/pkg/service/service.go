package service

import (
	"nats"
	cache "nats/pkg/cashe"
	"nats/pkg/repository"
)

type Creator interface {
	Create(ord nats.Order) error
	validate(ord nats.Order) error
}

type Reader interface {
	Restore() error
	Read(id string) (nats.OrderDTO, error)
}

type Service struct {
	Creator
	Reader
}

func NewService(r *repository.Repository, inMem cache.InMemory) *Service {
	return &Service{
		Creator: NewCreatorService(r.Creator, inMem),
		Reader:  NewReaderService(r.Reader, inMem),
	}
}
