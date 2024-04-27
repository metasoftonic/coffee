package models

import (
	"time"
)

type Extra struct {
	Id           int64     `db:"id"`
	Name         string    `db:"name"`
	Unit         int32     `db:"unit"`
	PricePerUnit float64   `db:"price_per_unit"`
	CreatedDate  time.Time `db:"created_date"`
}

type OrderExtra struct {
	Id           int64     `db:"id"`
	OrderId      string    `db:"order_id"`
	ExtraId      string    `db:"extra_id"`
	Quantity     int32     `db:"quantity"`
	PricePerUnit float64   `db:"price_per_unit"`
	CreatedDate  time.Time `db:"created_date"`
}
