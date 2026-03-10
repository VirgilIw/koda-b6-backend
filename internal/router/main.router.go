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

	RouterUser(app, c)
	RouterAuth(app, c)
	RouterProduct(app, c)
	RouterOrder(app, c)
}
