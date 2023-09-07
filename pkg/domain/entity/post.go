package entity

import (
	"time"
)

type Post struct {
	ID        int64    `bun:"id,pk,autoincrement"`
	Title     string
	Body      string   `bun:"type:longtext"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
	UserId    int64
}