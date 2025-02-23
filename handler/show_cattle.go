package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show cattle
// @Description Delete one cattle
// @Tags Cattle
// @Accept json
// @Produce json
// @Param id query string true "Cattle identification"
// @Success 200 {object} ShowCattleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cattle [get]
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
