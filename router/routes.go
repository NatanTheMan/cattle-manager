package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cattles", func (ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{"msg":"GET all cattles"})
    })
		v1.GET("/cattle", func (ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{"msg":"GET one cattle"})
    })
		v1.POST("/cattle", func (ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{"msg":"POST cattle"})
    })
    v1.PUT("/cattle/:id", func (ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{"msg":"PUT cattle"})
    })
		v1.DELETE("/cattle", func (ctx *gin.Context){
      ctx.JSON(http.StatusOK, gin.H{"msg":"DELETE cattle"})
    })
	}
}
