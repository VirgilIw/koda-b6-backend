package router

import (
	"github.com/gin-gonic/gin"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterUser(app *gin.Engine, c *di.Container) {

	user := app.Group("/users")

	handler := c.UserHandler()

	user.GET("", handler.GetUsers)
}
