package routes

import (
	"database/sql"
	"go-sanber64-quiz3/modules/books"
	"go-sanber64-quiz3/modules/categories"
	"go-sanber64-quiz3/modules/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	users.RegisterUserRoutes(api.Group("/users"), db)
	categories.RegisterCategoryRoutes(api.Group("/categories"), db)
	books.RegisterBookRoutes(api.Group("/books"), db)

	return router
}
