package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterProduct(app *gin.Engine, c *di.Container) {
	products := app.Group("/products")
	handler := c.ProductHandler()

	products.GET("", handler.GetProducts)
	products.GET("/:id", handler.GetProductById)
	products.POST("", handler.CreateProduct)
	products.PATCH("/:id", handler.UpdateProduct)
	products.DELETE("/:id", handler.DeleteProduct)
}
