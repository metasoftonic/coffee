package postgres

import (
	"fmt"
	"github/metasoftonic/coffee-api/internals/models"

	"github.com/jmoiron/sqlx"
)

func NewCoffeeRepository(db *sqlx.DB) *CoffeeRepository {
	return &CoffeeRepository{DB: db}
}

type CoffeeRepository struct {
	*sqlx.DB
}

func (c *CoffeeRepository) GetCoffee(id int64) (models.Coffee, error) {
	var t models.Coffee
	if err := c.Get(&t, `Select * from coffees where Id = $1`, id); err != nil {
		return models.Coffee{}, fmt.Errorf("error getting coffee : %w", err)
	}
	return t, nil
}

func (c *CoffeeRepository) GetCoffees() ([]models.Coffee, error) {
	var cc []models.Coffee
	if err := c.Select(&cc, `Select * from coffees`); err != nil {
		return []models.Coffee{}, fmt.Errorf("error getting coffees : %w", err)
	}
	return cc, nil
}

func (c *CoffeeRepository) CreateCoffee(coffee *models.Coffee) error {
	if err := c.Get(coffee, `Insert into coffees (name, description, size, price_per_unit, available_units) values ($1, $2, $3, $4, $5) returning *`,
		coffee.Name,
		coffee.Description,
		coffee.Size,
		coffee.PricePerUnit,
		coffee.AvailableUnits); err != nil {
		return fmt.Errorf("error inserting coffee : %w", err)
	}
	return nil
}

func (c *CoffeeRepository) UpdateCoffee(coffee *models.Coffee) error {
	if err := c.Get(coffee, `Update coffees set name = $1, description = $2, size = $3, price_per_unit = $4, available_units = $5) returning *`,
		coffee.Name,
		coffee.Description,
		coffee.Size,
		coffee.PricePerUnit,
		coffee.AvailableUnits); err != nil {
		return fmt.Errorf("error updating coffee : %w", err)
	}
	return nil
}
