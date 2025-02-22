package router

import (
	"github.com/NatanTheMan/cattle-manager/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
  handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cattles",handler.ListCattlesHandler)
		v1.GET("/cattle", handler.ShowCattleHandler)
		v1.POST("/cattle", handler.CreateCattleHandler)
    v1.PUT("/cattle/:id", handler.UpdateCattleHandler)
		v1.DELETE("/cattle", handler.DeleteCattleHandler)
	}
}
