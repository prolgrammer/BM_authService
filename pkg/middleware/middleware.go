package middleware

import "github.com/gin-gonic/gin"

type middleware struct {
}

type Middleware interface {
	HandleErrors(c *gin.Context)
}

func NewMiddleware() Middleware {
	return &middleware{}
}
