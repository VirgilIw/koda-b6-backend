package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterImages(app *gin.RouterGroup, c *di.Container) {
	images := app.Group("/images")
	handler := c.ImagesHandler()

	images.GET("", handler.GetImages)
	images.GET("/:id", handler.GetImageById)
	images.POST("", handler.CreateImage)
	images.PATCH("/:id", handler.UpdateImage)
	images.DELETE("/:id", handler.DeleteImageById)
}
