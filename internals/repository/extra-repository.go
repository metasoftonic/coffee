package repository

import (
	"github/metasoftonic/coffee-api/internals/models"
)

type ExtraRepository interface {
	GetExtra(id int64) (models.Extra, error)
	GetExtras() ([]models.Extra, error)
	CreateExtra(extra *models.Extra) error
	UpdateExtra(extra *models.Extra) error
}
