package router

import (
	docs "github.com/NatanTheMan/cattle-manager/docs"
	"github.com/NatanTheMan/cattle-manager/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/cattles", handler.ListCattlesHandler)
		v1.GET("/cattle", handler.ShowCattleHandler)
		v1.POST("/cattle", handler.CreateCattleHandler)
		v1.PUT("/cattle", handler.UpdateCattleHandler)
		v1.DELETE("/cattle", handler.DeleteCattleHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
