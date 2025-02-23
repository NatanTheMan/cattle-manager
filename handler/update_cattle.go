package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update cattle
// @Description Update a cattle
// @Tags Cattle
// @Accept json
// @Produce json
// @Param id query string true "Cattle identification"
// @Param cattle body UpdateCattleRequest true "Cattle data"
// @Success 200 {object} UpdateCattleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cattle [put]
func UpdateCattleHandler(ctx *gin.Context) {
	request := UpdateCattleRequest  {}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	cattle := schemas.Cattle{}
	if err := db.First(&cattle, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "cattle not found")
		return
	}

	if request.Name != "" {
		cattle.Name = request.Name
	}

	if err := db.Save(&cattle).Error; err != nil {
		logger.Errorf("error updating cattle: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating cattle")
		return
	}
	sendSuccess(ctx, "update-cattle", cattle)

}
