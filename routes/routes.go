package routes

import (
	"database/sql"
	"go-sanber64-quiz3/modules/books"
	"go-sanber64-quiz3/modules/categories"
	"go-sanber64-quiz3/modules/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	users.RegisterUserRoutes(api.Group("/users"), db)
	categories.RegisterCategoryRoutes(api.Group("/categories"), db)
	books.RegisterBookRoutes(api.Group("/books"), db)

	return router
}
