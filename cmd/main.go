package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load()
	port := os.Getenv("PORT")

	r.Run(fmt.Sprintf("localhost:%s", port))
}
