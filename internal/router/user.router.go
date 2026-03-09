package router

import (
	"github.com/gin-gonic/gin"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func RouterUser(app *gin.Engine, c *di.Container) {

	user := app.Group("/users", middleware.AuthMiddleware())

	handler := c.UserHandler()

	user.GET("", handler.GetUsers)
	user.PATCH("", handler.UpdateProfile)
}
