package dto

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	CategoryID  int    `json:"category_id"`
	ImageURL    string `json:"image_url"`
}