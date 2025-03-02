package main

import (
	"fmt"
	"go-sanber64-quiz3/databases/migrations"
	"go-sanber64-quiz3/routes"
	"go-sanber64-quiz3/utils"

	_ "go-sanber64-quiz3/docs"

	"time"

	"github.com/gin-contrib/cors"
	swaggFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Sanber64 Quiz3 API
// @version 1.0
// @description API untuk Books management
// @host https://go-sanber64-quiz3.onrender.com
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format "Bearer {token}"
func main() {
	DB, err := utils.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("❌ Error saat menginisialisasi database: %v", err))
	}
	defer DB.Close()

	migrations.Migrations(DB)

	router := routes.SetupRouter(DB)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggFiles.Handler))
	
	router.Run(":8080")
}
