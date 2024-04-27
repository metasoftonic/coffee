package repository

import (
	"github/metasoftonic/coffee-api/internals/models"
)

type CoffeeRepository interface {
	GetCoffee(id int64) (models.Coffee, error)
	GetCoffees() ([]models.Coffee, error)
	CreateCoffee(coffee *models.Coffee) error
	UpdateCoffee(coffee *models.Coffee) error
}
