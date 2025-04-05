package handlers

import (
	"github.com/gin-gonic/gin"
)

func respondWithJSON(ctx *gin.Context, statusCode int, responseData any) {
	ctx.JSON(statusCode, responseData)
}
