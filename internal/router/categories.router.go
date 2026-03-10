package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterCategories(app *gin.Engine, c *di.Container) {
	categories := app.Group("/categories")
	handler := c.CategoriesHandler()

	categories.GET("", handler.GetCategories)

}
