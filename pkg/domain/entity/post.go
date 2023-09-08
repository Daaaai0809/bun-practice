package entity

import (
	"time"
)

type Post struct {
	ID        int64    `bun:"id,pk,autoincrement"`
	Title     string
	Body      string   `bun:"type:longtext"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
	UserId    int64
}