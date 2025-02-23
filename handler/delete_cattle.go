package handler

import (
	"fmt"
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteCattleHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamter").Error())
		return
	}
	cattle := schemas.Cattle{}

	if err := db.First(&cattle, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("cattle with id: %s not found", id))
		return
	}

	if err := db.Delete(&cattle).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting cattle with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-cattle", cattle)
}
