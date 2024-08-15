package responses

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// type dataResponse struct {
// 	Data  interface{} `json:"data"`
// 	Count int64       `json:"count"`
// }

// type idResponse struct {
// 	ID interface{} `json:"id"`
// }

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(ctx *gin.Context, statusCode int, message string) {
	// logger.Error(message)
	code := strings.Split(message, ":")[0]
	ctx.AbortWithStatusJSON(statusCode, response{code, message})
}

func Data(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, struct {
		Data interface{} `json:"data"`
	}{Data: data})
}
