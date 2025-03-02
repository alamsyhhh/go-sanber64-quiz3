package main

import (
	"fmt"
	"go-sanber64-quiz3/config"
	"go-sanber64-quiz3/databases/migrations"
	"go-sanber64-quiz3/routes"

	_ "go-sanber64-quiz3/docs"

	swaggFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Sanber64 Quiz3 API
// @version 1.0
// @description API untuk user management di Go Gin
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format "Bearer {token}"
func main() {
	DB, err := config.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("❌ Error saat menginisialisasi database: %v", err))
	}
	defer DB.Close()

	migrations.Migrations(DB)

	router := routes.SetupRouter(DB)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggFiles.Handler))
	
	router.Run(":8080")
}
