package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterProduct(app *gin.RouterGroup, c *di.Container) {
	products := app.Group("/products")

	productHandler := c.ProductHandler()
	searchHandler := c.SearchHandler()

	products.GET("", productHandler.GetProducts)
	products.GET("/search", searchHandler.SearchProducts)
	products.GET("/:id", productHandler.GetDetailProductById)

	products.POST("", productHandler.CreateProduct)
	products.PATCH("/:id", productHandler.UpdateProduct)
	products.DELETE("/:id", productHandler.DeleteProduct)
}
