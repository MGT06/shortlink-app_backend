package model

import "time"

type Users struct {
	Id         int       `db:"id"`
	Name       string    `db:"full_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password_hash"`
	Created_at time.Time `db:"created_at"`
}
