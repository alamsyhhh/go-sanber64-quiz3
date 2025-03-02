package books

import (
	"go-sanber64-quiz3/modules/books/dto"
	"go-sanber64-quiz3/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service BookService
}

func NewBookController(service BookService) *BookController {
	return &BookController{service}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book with multipart form-data including an image
// @Tags Books
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Book Title"
// @Param description formData string true "Book Description"
// @Param release_year formData int true "Release Year"
// @Param price formData int true "Price"
// @Param total_page formData int true "Total Page"
// @Param category_id formData int true "Category ID"
// @Param image_url formData file true "Book Image"
// @Security BearerAuth
// @Router /api/books [post]
func (c *BookController) CreateBook(ctx *gin.Context) {
	var req dto.CreateBookRequest

	if err := ctx.ShouldBind(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdBy := ctx.GetString("username")

	file, err := ctx.FormFile("image_url")
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Gagal membaca gambar", nil)
		return
	}

	imageData, err := file.Open()
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Gagal membuka gambar", nil)
		return
	}
	defer imageData.Close()

	imageBytes := make([]byte, file.Size)
	_, err = imageData.Read(imageBytes)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Gagal membaca gambar", nil)
		return
	}

	book, err := c.service.CreateBook(req.Title, req.Description, imageBytes, req.ReleaseYear, req.Price, req.TotalPage, req.CategoryID, createdBy)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	responseData := gin.H{
		"id":          book.ID,
		"title":       book.Title,
		"description": book.Description,
		"image_url":   book.ImageURL,
		"release_year": book.ReleaseYear,
		"price":       book.Price,
		"total_page":  book.TotalPage,
		"thickness":   book.Thickness,
		"category_id": book.CategoryID,
		"created_at":  book.CreatedAt,
		"created_by":  book.CreatedBy,
		"modified_at": book.ModifiedAt,
		"modified_by": book.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "Buku berhasil ditambahkan", responseData)
}

// @Summary Get book by ID
// @Description Mendapatkan informasi buku berdasarkan ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "ID Buku"
// @Security BearerAuth
// @Router /api/books/{id} [get]
func (c *BookController) GetBookByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID harus berupa angka", nil)
		return
	}

	book, err := c.service.GetBookByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Detail buku ditemukan", book)
}

// @Summary Get all books
// @Description Mendapatkan daftar semua buku
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/books [get]
func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Daftar buku ditemukan", books)
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update book details with optional image update using multipart form-data
// @Tags Books
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Book ID"
// @Param title formData string true "Book Title"
// @Param description formData string true "Book Description"
// @Param release_year formData int true "Release Year"
// @Param price formData int true "Price"
// @Param total_page formData int true "Total Page"
// @Param category_id formData int true "Category ID"
// @Param image_url formData file false "Book Image (optional)"
// @Security BearerAuth
// @Router /api/books/{id} [put]
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID harus berupa angka", nil)
		return
	}

	var req dto.UpdateBookRequest
	if err := ctx.ShouldBind(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	modifiedBy := ctx.GetString("username")

	file, err := ctx.FormFile("image_url")
	var imageBytes []byte
	if err == nil {
		imageData, _ := file.Open()
		defer imageData.Close()
		imageBytes = make([]byte, file.Size)
		imageData.Read(imageBytes)
	}

	updatedBook, err := c.service.UpdateBook(id, req.Title, req.Description, imageBytes, req.ReleaseYear, req.Price, req.TotalPage, req.CategoryID, modifiedBy)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Buku berhasil diperbarui", updatedBook)
}

// @Summary Delete book by ID
// @Description Menghapus buku berdasarkan ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "ID Buku"
// @Security BearerAuth
// @Router /api/books/{id} [delete]
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID harus berupa angka", nil)
		return
	}

	deletedBook, err := c.service.DeleteBook(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Buku berhasil dihapus", deletedBook)
}
