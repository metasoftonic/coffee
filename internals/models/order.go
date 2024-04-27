package models

type Order struct {
	Id       int64 `db:"id"`
	CoffeeId int64 `db:"coffee_id"`
	UserId   int64 `db:"user_id"`
	Quantity int32 `db:"quantity"`
	Size     int32 `db:"size"`
	Extras   []*OrderExtra
}
