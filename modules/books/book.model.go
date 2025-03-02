package books

import "time"

type Book struct {
	ID          int       `db:"id,omitempty"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ImageURL    string    `db:"image_url"`
	ReleaseYear int       `db:"release_year"`
	Price       int       `db:"price"`
	TotalPage   int       `db:"total_page"`
	Thickness   string    `db:"thickness"`
	CategoryID  int       `db:"category_id"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   string    `db:"created_by"`
	ModifiedAt  time.Time `db:"modified_at"`
	ModifiedBy  string    `db:"modified_by"`
}
