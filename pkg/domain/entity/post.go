package entity

import (
	"time"
)

type Post struct {
	ID        int64    `bun:"id,pk,autoincrement"`
	Title     string
	Body      string   `bun:"type:longtext"`
	Tags      []*Tag   `bun:"m2m:post_tags,join:Post=Tag" json:"tags"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
	UserID    int64 `bun:"user_id"`
}