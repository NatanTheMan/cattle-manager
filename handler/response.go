package handler

import (
	"fmt"
	"net/http"

	"github.com/NatanTheMan/cattle-manager/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateCattleResponse struct {
	Message string                 `json:"message"`
	Data    schemas.CattleResponse `json:"data"`
}

type DeleteCattleResponse struct {
	Message string                 `json:"message"`
	Data    schemas.CattleResponse `json:"data"`
}

type ShowCattleResponse struct {
	Message string                 `json:"message"`
	Data    schemas.CattleResponse `json:"data"`
}

type ListCattlesResponse struct {
	Message string                   `json:"message"`
	Data    []schemas.CattleResponse `json:"data"`
}

type UpdateCattleResponse struct {
	Message string                 `json:"message"`
	Data    schemas.CattleResponse `json:"data"`
}
