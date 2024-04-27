package models

import "time"

type User struct {
	Id          int64     `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	CreatedDate time.Time `db:"created_date"`
}
