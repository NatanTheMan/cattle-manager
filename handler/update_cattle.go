package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCattleHandler(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"msg":"sdlfkjsakdj"})
}
