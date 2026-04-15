package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func RouterProduct(app *gin.RouterGroup, c *di.Container) {
	products := app.Group("/products")

	productHandler := c.ProductHandler()
	searchHandler := c.SearchHandler()

	products.GET("", productHandler.GetProducts)
	products.GET("/search", searchHandler.SearchProducts)
	products.GET("/:id", productHandler.GetDetailProductById)

	protected := products.Group("")
	protected.Use(middleware.AuthMiddleware(), middleware.RBACMiddleware())

	protected.POST("", productHandler.CreateProduct)
	protected.PATCH("/:id", productHandler.UpdateProduct)
	protected.DELETE("/:id", productHandler.DeleteProduct)
}
