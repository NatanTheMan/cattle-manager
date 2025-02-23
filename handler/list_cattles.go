package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List cattles
// @Description List all cattles
// @Tags Cattle
// @Accept json
// @Produce json
// @Success 200 {object} ListCattlesResponse
// @Failure 500 {object} ErrorResponse
// @Router /cattles [get]
func ListCattlesHandler(ctx *gin.Context) {
	cattles := []schemas.Cattle{}
	if err := db.Find(&cattles).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing cattles")
		return
	}

	sendSuccess(ctx, "list-cattle", cattles)
}
