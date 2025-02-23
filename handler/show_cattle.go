package handler

import (
	"fmt"
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show cattle
// @Description Show one cattle
// @Tags Cattle
// @Accept json
// @Produce json
// @Param id query string false "Cattle identification"
// @Param earring query string false "Cattle earring"
// @Success 200 {object} ShowCattleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cattle [get]
func ShowCattleHandler(ctx *gin.Context) {
	cattle := schemas.Cattle{}
	id := ctx.Query("id")
	earring  := ctx.Query("earring")

	if id == "" && earring == "" {
		sendError(ctx, http.StatusBadRequest, "one must param")
		return
	}
	if id != "" {
		if err := db.First(&cattle, id).Error; err != nil {
			sendError(ctx, http.StatusNotFound, "cattle not found")
			return
		}
	}
	if earring != "" {
		if err := db.Find(&cattle,"earring = ?", earring).Error; err != nil {
			sendError(ctx, http.StatusNotFound, "cattle not found")
			return
		}
	}

	sendSuccess(ctx, "show-cattle", cattle)
}
