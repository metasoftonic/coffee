package repository

type Store interface {
	CoffeeRepository
	ExtraRepository
	OrderRepository
}
