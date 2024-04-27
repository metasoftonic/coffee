package postgres

import (
	"fmt"
	"github/metasoftonic/coffee-api/internals/models"

	"github.com/jmoiron/sqlx"
)

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

type OrderRepository struct {
	*sqlx.DB
}

func (c *OrderRepository) GetOrder(id int64) (models.Order, error) {
	var t models.Order
	if err := c.Get(&t, `Select * from orders where Id = $1`, id); err != nil {
		return models.Order{}, fmt.Errorf("error getting ordera : %w", err)
	}
	return t, nil
}

func (c *OrderRepository) GetOrders() ([]models.Order, error) {
	var cc []models.Order
	if err := c.Select(&cc, `Select * from orders`); err != nil {
		return []models.Order{}, fmt.Errorf("error getting orders : %w", err)
	}
	return cc, nil
}

func (c *OrderRepository) CreateOrder(order *models.Order) error {
	tx, err := c.Beginx()
	if err != nil {
		return fmt.Errorf("error creating orders : %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	query := `Insert into orders (coffee_id, user_id, quantity, size) values(?,?,?)`

	result, err := tx.Exec(query, order.UserId, order.CoffeeId, order.Quantity, order.Size)
	if err != nil {
		return fmt.Errorf("error creating orders : %w", err)
	}
	orderId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error creating order created record : %w", err)
	}

	for _, extra := range order.Extras {

		//get the extra,
		var e models.Extra
		if err = tx.Get(&e, `Select id, price_per_unit from extras where id = $1`, extra.ExtraId); err != nil {
			return fmt.Errorf("error fetching extra for order while creating : %w", err)
		}
		query = `
            INSERT INTO order_extras (order_id, extra_id, quantity, price_per_unit)
            VALUES (?, ?, ?, ?)
        `
		_, err := tx.Exec(query, orderId, extra.ExtraId, extra.Quantity, extra.PricePerUnit)
		if err != nil {
			return fmt.Errorf("error creating order created record : %w", err)
		}
	}
	return nil
}
func (c *OrderRepository) UpdateOrder(coffee *models.Order) error {
	panic("not implemented")
}
