package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prolgrammer/BM_authService/controllers"
	"net/http"
)

func (m *middleware) Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := m.manager.Parse(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.ErrAuthenticated)
		return
	}

	c.Set("account_id", claims["sub"])
	c.Next()
}
