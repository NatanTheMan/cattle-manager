package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create cattle
// @Description Create a new cattle
// @Tags Cattle
// @Accept json
// @Produce json
// @Param request body CreateCattleRequest true "Request body"
// @Success 200 {object} CreateCattleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cattle [post]
func CreateCattleHandler(ctx *gin.Context) {
  cattle := schemas.Cattle{}
	request := CreateCattleRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
  if err:=db.Find(&cattle, "earring = ?", request.Earring).Error; err == nil {
    sendError(ctx,http.StatusBadRequest, fmt.Sprintf("a cattle with earring: %s already exists", request.Earring))
    return
  }

	cattle = schemas.Cattle{
		Name:    request.Name,
		Earring: request.Earring,
		Gender:  strings.ToUpper(request.Gender),
	}

	if err := db.Create(&cattle).Error; err != nil {
		logger.Errorf("error creating cattle: %v", err)
		sendError(ctx, http.StatusBadRequest, "error creating cattle on database")
		return
	}

	sendSuccess(ctx, "create-cattle", cattle)
}
