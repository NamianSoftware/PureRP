package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ResponseError struct {
	Message string `json:"message"`
}

func newResponseError(c *gin.Context, statusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(statusCode, ResponseError{message})
}
