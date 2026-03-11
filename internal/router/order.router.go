package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterOrder(app *gin.RouterGroup, c *di.Container) {
	orders := app.Group("/coupons")
	handler := c.OrderHandler()

	orders.GET("/:id", handler.GetCouponById)
	orders.GET("", handler.GetCoupons)
	orders.PATCH("/:id", handler.EditCoupon)
	orders.POST("", handler.CreateCoupon)
	orders.DELETE("/:id", handler.DeleteCoupon)
}
