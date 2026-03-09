package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterAuth(app *gin.Engine, c *di.Container) {
	auth := app.Group("/auth")
	handler := c.AuthHandler()

	auth.POST("/login", handler.AuthLogin)
}
