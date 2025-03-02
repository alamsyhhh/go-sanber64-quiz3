package users

import (
	"time"
)

type User struct {
	ID         int       `db:"id,omitempty"`
	Username   string    `db:"username"`
	Password   string    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
	CreatedBy  string    `db:"created_by"`
	ModifiedAt time.Time `db:"modified_at"`
	ModifiedBy string    `db:"modified_by"`
}
