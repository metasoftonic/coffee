package postgres

import (
	"fmt"
	"github/metasoftonic/coffee-api/internals/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(datasource string) (*Store, error) {
	db, err := sqlx.Open("postgres", datasource)
	if err != nil {
		return nil, fmt.Errorf("error opening database :%w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database :%w", err)
	}
	return &Store{
		CoffeeRepository: NewCoffeeRepository(db),
		ExtraRepository:  NewExtraRepository(db),
		OrderRepository:  NewOrderRepository(db),
	}, nil
}

type Store struct {
	repository.CoffeeRepository
	repository.ExtraRepository
	repository.OrderRepository
}
