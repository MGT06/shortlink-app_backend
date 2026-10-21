package model

import "time"

type Links struct {
	LinkId       int       `db:"id"`
	UserId       int       `db:"user_id"`
	OriginaluUrl string    `db:"original_url"`
	Slug         string    `db:"slug"`
	Deleted_at   time.Time `db:"delete_at"`
	Created_at   time.Time `db:"created_at"`
}
