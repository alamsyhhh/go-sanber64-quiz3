package dto

type CreateBookRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	ReleaseYear int    `form:"release_year" binding:"required"`
	Price       int    `form:"price" binding:"required"`
	TotalPage   int    `form:"total_page" binding:"required"`
	CategoryID  int    `form:"category_id" binding:"required"`
}

type UpdateBookRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	ReleaseYear int    `form:"release_year" binding:"required"`
	Price       int    `form:"price" binding:"required"`
	TotalPage   int    `form:"total_page" binding:"required"`
	CategoryID  int    `form:"category_id" binding:"required"`
}
