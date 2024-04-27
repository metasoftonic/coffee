package postgres

import (
	"fmt"
	"github/metasoftonic/coffee-api/internals/models"

	"github.com/jmoiron/sqlx"
)

func NewExtraRepository(db *sqlx.DB) *ExtraRepository {
	return &ExtraRepository{DB: db}
}

type ExtraRepository struct {
	*sqlx.DB
}

func (c *ExtraRepository) GetExtra(id int64) (models.Extra, error) {
	var t models.Extra
	if err := c.Get(&t, `Select * from extras where Id = $1`, id); err != nil {
		return models.Extra{}, fmt.Errorf("error getting extras : %w", err)
	}
	return t, nil
}

func (c *ExtraRepository) GetExtras() ([]models.Extra, error) {
	var cc []models.Extra
	if err := c.Select(&cc, `Select * from extras`); err != nil {
		return []models.Extra{}, fmt.Errorf("error getting extras : %w", err)
	}
	return cc, nil
}

func (c *ExtraRepository) CreateExtra(e *models.Extra) error {
	if err := c.Get(e, `Insert into extras values ($1, $2, $3) returning *`,
		e.Name,
		e.Unit,
		e.PricePerUnit); err != nil {
		return fmt.Errorf("error inserting extras : %w", err)
	}
	return nil
}

func (c *ExtraRepository) UpdateExtra(e *models.Extra) error {

	if err := c.Get(e, `Update extras set name = $1, unit = $2, price_per_unit = $4) returning *`,
		e.Name,
		e.Unit,
		e.PricePerUnit); err != nil {
		return fmt.Errorf("error updating extras : %w", err)
	}
	return nil
}
