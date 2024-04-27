package repository

import (
	"github/metasoftonic/coffee-api/internals/models"
)

type OrderRepository interface {
	GetOrder(id int64) (models.Order, error)
	GetOrders() ([]models.Order, error)
	CreateOrder(extra *models.Order) error
	UpdateOrder(extra *models.Order) error
}
