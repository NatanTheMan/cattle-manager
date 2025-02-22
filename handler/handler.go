package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCattleHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}

func ShowCattleHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}

func UpdateCattleHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}

func DeleteCattleHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}

func ListCattlesHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}
