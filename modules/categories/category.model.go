package categories

import (
	"time"
)

type Category struct {
	ID         int       `db:"id,omitempty"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	CreatedBy  string    `db:"created_by"`
	ModifiedAt time.Time `db:"modified_at"`
	ModifiedBy string    `db:"modified_by"`
}