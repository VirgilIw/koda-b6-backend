package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterSizes(app *gin.RouterGroup, c *di.Container) {
	sizes := app.Group("/sizes")
	handler := c.SizesHandler()

	sizes.GET("", handler.GetSizes)
	sizes.GET("/:id", handler.GetSizeByID)
	sizes.PATCH("/:id", handler.UpdateSize)
	sizes.DELETE("/:id", handler.DeleteSizeById)

}
