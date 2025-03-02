package users

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"

	"go-sanber64-quiz3/middlewares"
)


func RegisterUserRoutes(router *gin.RouterGroup, db *sql.DB) {
	// Konversi *sql.DB ke *goqu.Database
	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)

	// Inisialisasi repository, service, dan controller
	userRepo := NewUserRepository(goquDB)
	userService := NewUserService(userRepo)
	userController := NewUserController(userService)

	// Rute untuk register dan login
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Rute yang membutuhkan autentikasi JWT
	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.PUT("/update", userController.UpdateUser)
	protected.GET("/me", userController.GetMe)
}

