package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/config"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/router"
)

// @title Coffeshop Backend
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
		log.Fatal(err)
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rdb := config.InitRedis()

	app := gin.Default()

	container := di.NewContainer(db, rdb)

	router.Init(app, container)

	port := os.Getenv("PORT")

	app.Run(fmt.Sprintf("localhost:%s", port))
}
