package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

func ShowCattleHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	cattle := schemas.Cattle{}
	if err := db.First(&cattle,id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "cattle not found")
    return
	}

	sendSuccess(ctx, "show-cattle", cattle)
}
