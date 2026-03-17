package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virgiIw/koda-b6-coffeshopdb/docs"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/config"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/router"
)

func main() {
	fmt.Println("backend testt")

	_ = godotenv.Load()

	docs.SwaggerInfo.Title = "Coffeshop Backend"
	docs.SwaggerInfo.Description = "Coffeshop BE documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("APP_URL")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("failed connect database: %v", err)
	}
	defer db.Close()

	rdb := config.InitRedis()

	app := gin.Default()

	container := di.NewContainer(db, rdb)

	router.Init(app, container)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	app.Run(":" + port)
}
