package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, data any, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

func InternalServerError(c *gin.Context) {
	Response(c, nil, http.StatusInternalServerError, "internal server error")
}

func NotFound(c *gin.Context) {
	Response(c, nil, http.StatusNotFound, "not found")
}
