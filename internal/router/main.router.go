package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/virgiIw/koda-b6-coffeshopdb/internal/middleware"
)

func Init(app *gin.Engine, db *pgx.Conn) {

	app.Use(middleware.CorsMiddleware())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
