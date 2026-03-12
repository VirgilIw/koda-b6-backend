package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
)

func RouterLanding(app *gin.Engine, c *di.Container) {
	reviews := app.Group("/reviews")
	handler := c.LandingHandler()

	reviews.GET("", handler.GetReviews)
}
