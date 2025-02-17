package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prolgrammer/BM_authService/controllers"
	"github.com/prolgrammer/BM_authService/internal/repositories"
	"github.com/prolgrammer/BM_authService/internal/usecases"
	"net/http"
)

func (m middleware) HandleErrors(c *gin.Context) {
	if len(c.Errors) > 0 {
		err := c.Errors.Last()

		fmt.Println("an error has happened")
		if errors.Is(err, controllers.ErrDataBindError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, controllers.ErrRegistrationsError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, controllers.ErrAuthRequiredError) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		//repositories

		if errors.Is(err, repositories.ErrEntityNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}

		//UseCases

		if errors.Is(err, usecases.ErrEntityAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusConflict, err.Error())
			return
		}

		if errors.Is(err, usecases.ErrPasswordMismatch) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		fmt.Println("Unexpected error")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
