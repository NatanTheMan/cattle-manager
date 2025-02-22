package handler

import (
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCattleHandler(ctx *gin.Context) {
	request := CreateCattleRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

  cattle := schemas.Cattle{
    Name:request.Name,
    Earring:request.Earring,
    Gender:request.Gender,
  }

	if err := db.Create(&cattle).Error; err != nil {
		logger.Errorf("error creating cattle: %v", err)
		sendError(ctx, http.StatusBadRequest, "error creating cattle on database")
		return
	}

  sendSuccess(ctx,  "create-cattle", cattle)
}
