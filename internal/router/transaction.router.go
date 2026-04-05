package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterTransactions(app *gin.RouterGroup, c *di.Container) {
	transactions := app.Group("/transactions")
	handler := c.TransactionHandler()

	// transactions.GET("", handler.GetTransactions)
	// transactions.GET("/:id", handler.GetTransactionDetail)
	transactions.POST("", handler.CreateTransaction)
	transactions.PATCH("/:id/status", handler.UpdateTransactionStatus)
	// transactions.DELETE("/:id", handler.DeleteTransactionById)
}
