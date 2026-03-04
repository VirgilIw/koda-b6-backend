package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/config"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/router"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	db, err := config.InitDB()

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	defer db.Close(context.Background())

	router.Init(r, db)

	port := os.Getenv("PORT")

	r.Run("localhost:" + port)
}
