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

type signupController struct {
	user usecases.SignUpUseCase
}

func NewSignUpController(
	handler *gin.Engine,
	user usecases.SignUpUseCase,
	middleware middleware.Middleware,
) {
	u := &signupController{
		user: user,
	}

	handler.POST("/signup", u.SignUp, middleware.HandleErrors)
}

func (u *signupController) SignUp(ctx *gin.Context) {
	fmt.Print("SignUp\n")

	var req requests.SignRequest
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		wrappedErr := fmt.Errorf("%w: %v", controllers.ErrDataBindError, err)
		middleware.AddGinError(ctx, wrappedErr)

		return
	}

	response, err := u.user.SignUp(ctx, req)
	if err != nil {
		middleware.AddGinError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
