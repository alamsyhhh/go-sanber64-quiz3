package books

import (
	"database/sql"
	"go-sanber64-quiz3/middlewares"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.RouterGroup, db *sql.DB) {
	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)

	repo := NewBookRepository(goquDB)
	svc := NewBookService(repo)
	ctrl := NewBookController(svc)

	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.POST("/", ctrl.CreateBook)
	protected.GET("/", ctrl.GetAllBooks)
	protected.GET("/:id", ctrl.GetBookByID)
	protected.PUT("/:id", ctrl.UpdateBook)
	protected.DELETE("/:id", ctrl.DeleteBook)
}
