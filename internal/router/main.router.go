package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/virgiIw/koda-b6-coffeshopdb/docs"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/di"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func Init(app *gin.Engine, c *di.Container) {

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Use(middleware.CorsMiddleware())

	// auth tetap public
	RouterAuth(app, c)
	RouterRecommendProduct(app, c)
	RouterLanding(app, c)
	RouterCart(app, c)
	// admin routes
	admin := app.Group("/admin")

	RouterUser(admin, c)
	RouterProduct(admin, c)
	RouterOrder(admin, c)
	RouterCategories(admin, c)
	RouterSizes(admin, c)
	RouterVariant(admin, c)
	RouterImages(admin, c)
}
