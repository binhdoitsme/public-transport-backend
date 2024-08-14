package responses

import "github.com/gin-gonic/gin"

// type dataResponse struct {
// 	Data  interface{} `json:"data"`
// 	Count int64       `json:"count"`
// }

// type idResponse struct {
// 	ID interface{} `json:"id"`
// }

type response struct {
	Message string `json:"message"`
}

func Error(ctx *gin.Context, statusCode int, message string) {
	// logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{message})
}
