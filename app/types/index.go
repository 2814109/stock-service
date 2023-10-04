package types

import (
	"time"
)

type Post struct {
	ID        int64 `bun:",pk,autoincrement"`
	Content   string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
