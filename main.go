package main

import (
	"fmt"
	"go-sanber64-quiz3/databases/migrations"
	"go-sanber64-quiz3/routes"
	"go-sanber64-quiz3/utils"

	_ "go-sanber64-quiz3/docs"

	swaggFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Sanber64 Quiz3 API
// @version 1.0
// @description API untuk Books management
// @host go-sanber64-quiz3.onrender.com
// @BasePath /
// @securityDefinitions.apikey BearerAuth
func main() {
	DB, err := utils.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("❌ Error saat menginisialisasi database: %v", err))
	}
	defer DB.Close()

	migrations.Migrations(DB)

	router := routes.SetupRouter(DB)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggFiles.Handler))
	
	router.Run(":8080")
}
