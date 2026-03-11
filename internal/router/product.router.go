package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterProduct(app *gin.RouterGroup, c *di.Container) {
	products := app.Group("/products")
	handler := c.ProductHandler()

	products.GET("", handler.GetProducts)
	products.GET("/:id", handler.GetDetailProductById)
	products.POST("", handler.CreateProduct)
	products.PATCH("/:id", handler.UpdateProduct)
	products.DELETE("/:id", handler.DeleteProduct)
}
