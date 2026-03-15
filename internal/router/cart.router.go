package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func RouterCart(app *gin.Engine, c *di.Container) {

	handler := c.CartHandler()

	cart := app.Group("/cart")
	cart.Use(middleware.AuthMiddleware())

	cart.POST("", handler.AddToCart)
	cart.GET("", handler.GetCart)
	// cart.PATCH("/:id", handler.UpdateQty)
	// cart.DELETE("/:id", handler.DeleteCart)
}
