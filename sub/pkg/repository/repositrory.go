package repository

import (
	"github.com/jmoiron/sqlx"
	"nats"
)

type Creator interface {
	Create(ord *nats.Order) error
}

type Reader interface {
	Recover() ([]nats.Order, error)
	getItems(id string) ([]nats.Item, error)
	getDel(id string) (nats.Del, error)
	getPayment(id string) (nats.Pay, error)
}

type Repository struct {
	Creator
	Reader
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Creator: NewCreatorPostgres(db),
		Reader:  NewReaderPostgres(db),
	}
}
