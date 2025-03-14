package categories

import (
	"go-sanber64-quiz3/modules/categories/dto"
	"go-sanber64-quiz3/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{service}
}

// CreateCategory godoc
// @Summary Membuat kategori baru
// @Description Membuat kategori baru dengan nama yang diberikan
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body dto.CreateCategoryRequest true "Data kategori"
// @Security BearerAuth
// @Router /api/categories [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdBy, exists := ctx.Get("username")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	category, err := c.service.CreateCategory(req.Name, createdBy.(string))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	responseData := gin.H{
		"id":         category.ID,
		"name":       category.Name,
		"created_at": category.CreatedAt,
		"created_by": category.CreatedBy,
		"modified_at": category.ModifiedAt,
		"modified_by": category.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "Kategori berhasil dibuat", responseData)
}

// GetAllCategories godoc
// @Summary Mendapatkan semua kategori
// @Description Mengembalikan daftar semua kategori yang tersedia
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Router /api/categories [get]
func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Daftar kategori ditemukan", categories)
}

// GetCategoryByID godoc
// @Summary Mendapatkan detail kategori berdasarkan ID
// @Description Mengambil informasi kategori berdasarkan ID yang diberikan
// @Tags Categories
// @Produce json
// @Param id path int true "ID Kategori"
// @Security BearerAuth
// @Router /api/categories/{id} [get]
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID kategori tidak valid", nil)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	modifiedBy, exists := ctx.Get("username")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	category, err := c.service.UpdateCategory(id, req.Name, modifiedBy.(string))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	responseData := gin.H{
		"id":          category.ID,
		"name":        category.Name,
		"created_at":  category.CreatedAt,
		"created_by":  category.CreatedBy,
		"modified_at": category.ModifiedAt,
		"modified_by": category.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "Kategori berhasil diperbarui", responseData)
}


// UpdateCategory godoc
// @Summary Memperbarui kategori berdasarkan ID
// @Description Mengupdate nama kategori yang telah ada
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "ID Kategori"
// @Param request body dto.UpdateCategoryRequest true "Data kategori"
// @Security BearerAuth
// @Router /api/categories/{id} [put]
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID kategori tidak valid", nil)
		return
	}

	category, err := c.service.DeleteCategory(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	responseData := gin.H{
		"id":          category.ID,
		"name":        category.Name,
		"created_at":  category.CreatedAt,
		"created_by":  category.CreatedBy,
		"modified_at": category.ModifiedAt,
		"modified_by": category.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "Kategori berhasil dihapus", responseData)
}


// DeleteCategory godoc
// @Summary Menghapus kategori berdasarkan ID
// @Description Menghapus kategori dengan ID tertentu
// @Tags Categories
// @Produce json
// @Param id path int true "ID Kategori"
// @Security BearerAuth
// @Router /api/categories/{id} [delete]
func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID kategori tidak valid", nil)
		return
	}

	category, err := c.service.GetCategoryByID(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Kategori tidak ditemukan", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Detail kategori ditemukan", category)
}

// GetBooksByCategory godoc
// @Summary Mendapatkan daftar buku berdasarkan kategori
// @Description Mengambil daftar buku yang tersedia berdasarkan kategori tertentu
// @Tags Categories
// @Produce json
// @Param id path int true "ID Kategori"
// @Security BearerAuth
// @Router /api/categories/{id}/books [get]
func (c *CategoryController) GetBooksByCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "ID kategori tidak valid", nil)
		return
	}

	books, err := c.service.GetBooksByCategory(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusNotFound, "Kategori tidak ditemukan atau tidak memiliki buku", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Daftar buku ditemukan", books)
}
