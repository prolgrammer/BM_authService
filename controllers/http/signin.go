package http

import (
	"auth/controllers"
	"auth/controllers/http/middleware"
	"auth/controllers/requests"
	"auth/internal/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInController struct {
	user usecases.SignInUseCase
}

func NewSignInController(
	handler *gin.Engine,
	signInUseCase usecases.SignInUseCase,
	middleware middleware.Middleware,
) {

	u := &signInController{
		user: signInUseCase,
	}

	handler.POST("/signin", u.SignIn, middleware.HandleErrors)
}

func (uc *signInController) SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var request requests.SignRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		wrappedError := fmt.Errorf("%w: %w", controllers.ErrDataBindError, err)
		middleware.AddGinError(ctx, wrappedError)
		return
	}

	response, err := uc.user.SignIn(ctx, request)
	if err != nil {
		middleware.AddGinError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
