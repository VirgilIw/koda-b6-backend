package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func RouterTransactions(app *gin.Engine, c *di.Container) {
	transactions := app.Group("/transactions", middleware.AuthMiddleware())

	h := c.TransactionHandler()

	transactions.POST("/", h.CreateTransaction)
	transactions.GET("/", h.GetTransactionsByUserID)
	transactions.GET("/:id", h.GetTransactionDetail)
}

func RouterAdminTransactions(app *gin.RouterGroup, c *di.Container) {
	transactions := app.Group("/transactions", middleware.AuthMiddleware(), middleware.RBACMiddleware())

	h := c.TransactionHandler()

	transactions.PATCH("/:id/status", h.UpdateTransactionStatus)

}
