package entity

import (
	"time"
)

type User struct {
	ID        int64    `bun:"id,pk,autoincrement"`
	Email     string   `bun:"type:varchar(255),unique"`
	Password  string   `bun:"type:varchar(255)"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
	Posts	  []*Post	`bun:"rel:has-many,join:id=user_id"`
}