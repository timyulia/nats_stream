package service

import (
	"encoding/json"
	"nats"
	cache "nats/pkg/cashe"
	"nats/pkg/repository"
)

type ReaderService struct {
	repo  repository.Reader
	inMem cache.InMemory
}

func NewReaderService(repo repository.Reader, inMem cache.InMemory) *ReaderService {
	return &ReaderService{repo: repo, inMem: inMem}
}

func (s *ReaderService) Restore() error {
	ords, err := s.repo.Recover()
	if err != nil {
		return err
	}
	return s.inMem.RestoreCache(ords)
}

func (s *ReaderService) Read(id string) (nats.OrderDTO, error) {
	ord, err := s.inMem.ReadOrder(id)
	if err != nil {
		return nats.OrderDTO{}, err
	}
	ordJSON, err := json.Marshal(ord)
	if err != nil {
		return nats.OrderDTO{}, err
	}
	var ordDTO nats.OrderDTO
	err = json.Unmarshal(ordJSON, &ordDTO)
	return ordDTO, err
}
