package models

import (
	"time"
)

type Coffee struct {
	Id             int64     `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Size           float64   `db:"size"`
	PricePerUnit   float64   `db:"price_per_unit"`
	AvailableUnits int64     `db:"available_units"`
	CreatedDate    time.Time `db:"created_date"`
}
