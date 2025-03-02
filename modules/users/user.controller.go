package users

import (
	"go-sanber64-quiz3/modules/users/dto"
	"go-sanber64-quiz3/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service}
}

// Register godoc
// @Summary Mendaftarkan user baru
// @Description Endpoint untuk registrasi user baru
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Data User"
// @Success 200 {object} map[string]interface{} "User berhasil didaftarkan"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /api/users/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.RegisterUser(req.Username, req.Password)

	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", gin.H{"error": err.Error()})
		return
	}

	responseData := gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": user.CreatedAt,
		"created_by": user.CreatedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "User berhasil didaftarkan", responseData)
}

// Login godoc
// @Summary Login user
// @Description Endpoint untuk login user dan mendapatkan token JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Data Login"
// @Success 200 {object} map[string]interface{} "Login berhasil"
// @Failure 401 {object} map[string]interface{} "Username atau password salah"
// @Router /api/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.LoginUser(req.Username, req.Password)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Login berhasil", gin.H{"token": token})
}

// UpdateUser godoc
// @Summary Update user
// @Description Endpoint untuk memperbarui informasi user
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserRequest true "Data User"
// @Success 200 {object} map[string]interface{} "User berhasil diperbarui"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/users/update [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.service.UpdateUser(userID.(int), req.Username)
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", gin.H{"error": err.Error()})
		return
	}

	responseData := gin.H{
		"id":          updatedUser.ID,
		"username":    updatedUser.Username,
		"created_at":  updatedUser.CreatedAt,
		"created_by":  updatedUser.CreatedBy,
		"modified_at": updatedUser.ModifiedAt,
		"modified_by": updatedUser.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "User berhasil diperbarui", responseData)
}

// GetMe godoc
// @Summary Mendapatkan data user yang sedang login
// @Description Endpoint untuk mendapatkan informasi user berdasarkan token
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "User ditemukan"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/users/me [get]
func (c *UserController) GetMe(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "User not found", nil)
		return
	}

	user, err := c.service.GetMe(userID.(int))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user", gin.H{"error": err.Error()})
		return
	}

	responseData := gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"created_at":  user.CreatedAt,
		"created_by":  user.CreatedBy,
		"modified_at": user.ModifiedAt,
		"modified_by": user.ModifiedBy,
	}

	common.GenerateSuccessResponseWithData(ctx, "User ditemukan", responseData)
}


