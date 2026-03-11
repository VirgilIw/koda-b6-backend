package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterRecommendProduct(app *gin.Engine, c *di.Container) {
	recommProducts := app.Group("/recommended-products")

	handler := c.ProductHandler()

	recommProducts.GET("", handler.GetRecommendedProducts)
}
