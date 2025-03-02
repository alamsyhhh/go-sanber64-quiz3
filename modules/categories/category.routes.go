package categories

import (
	"database/sql"
	"go-sanber64-quiz3/middlewares"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.RouterGroup, db *sql.DB) {
	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)

	repo := NewCategoryRepository(goquDB)
	svc := NewCategoryService(repo)
	ctrl := NewCategoryController(svc)

	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.POST("/", ctrl.CreateCategory)
	protected.GET("/", ctrl.GetAllCategories)
	protected.PUT("/:id", ctrl.UpdateCategory)
	protected.DELETE("/:id", ctrl.DeleteCategory)
	protected.GET("/:id", ctrl.GetCategoryByID)
	protected.GET("/:id/books", ctrl.GetBooksByCategory)

}
