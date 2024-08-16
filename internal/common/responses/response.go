package responses

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

func Error(ctx *gin.Context, statusCode int, message string) {
	// logger.Error(message)
	code := strings.Split(message, ":")[0]
	ctx.AbortWithStatusJSON(statusCode, response{code, message})
}

func Data(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.AbortWithStatusJSON(statusCode, &dataResponse{Data: data})
}
