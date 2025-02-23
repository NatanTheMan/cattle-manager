package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

func ListCattlesHandler(ctx *gin.Context) {
	cattles := []schemas.Cattle{}
	if err := db.Find(&cattles).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing cattles")
		return
	}

	sendSuccess(ctx, "list-cattle", cattles)
}
