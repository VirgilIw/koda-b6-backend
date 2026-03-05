package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virgiIw/koda-b6-coffeshopdb/docs"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/config"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/router"
)

// @title Coffeshop BackendSELECT * FROM products
// @version 1.0
// @description Coffeshop BE documentation
// @host localhost:8888
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token. Example: "Bearer eyJhbGciO..."
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	docs.SwaggerInfo.BasePath = "/"

	r := gin.Default()

	db, err := config.InitDB()

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	defer db.Close(context.Background())

	container := di.NewContainer(db)

	router.Init(r, container)

	port := os.Getenv("PORT")

	r.Run("localhost:" + port)
}
