package entity

import (
	"time"
)

type Tag struct {
	ID	int64 `bun:"id,pk,autoincrement" json:"id"`
	Name	string `bun:"name" json:"name"`
	Posts	[]*Post `bun:"m2m:post_tags,join:Tag=Post" json:"posts"` 
	CreatedAt	time.Time `bun:"created_at" json:"created_at"`
	UpdatedAt	time.Time `bun:"updated_at" json:"updated_at"`
}