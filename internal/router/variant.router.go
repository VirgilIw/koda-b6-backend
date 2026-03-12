package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterVariant(app *gin.RouterGroup, c *di.Container) {
	variants := app.Group("/variants")
	handler := c.VariantHandler()

	variants.GET("", handler.GetVariants)
	variants.GET("/:id", handler.GetVariantById)
	variants.DELETE("/:id", handler.DeleteVariant)
}
