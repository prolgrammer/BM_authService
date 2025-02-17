package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prolgrammer/BM_authService/controllers"
	"github.com/prolgrammer/BM_authService/controllers/requests"
	"github.com/prolgrammer/BM_authService/internal/usecases"
	middleware2 "github.com/prolgrammer/BM_authService/pkg/middleware"
	"net/http"
)

type signupController struct {
	user usecases.SignUpUseCase
}

func NewSignUpController(
	handler *gin.Engine,
	user usecases.SignUpUseCase,
	middleware middleware2.Middleware,
) {
	u := &signupController{
		user: user,
	}

	handler.POST("/signup", u.SignUp, middleware.HandleErrors)
}

// SignUp godoc
// @Summary регистрация пользователя
// @Description регистрация пользователя в систему
// @Accept json
// @Produce json
// @Param request body requests.SignRequest true "структура запроса"
// @Success 200 {object} responses.SignResponse
// @Failure 400 {object} string "некорректный формат запроса"
// @Failure 409 {object} string "пользователь уже существует"
// @Failure 500 {object} string "внутренняя ошибка сервера"
// @Router /signup [post]
func (u *signupController) SignUp(ctx *gin.Context) {
	fmt.Print("SignUp\n")

	var req requests.SignRequest
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		wrappedErr := fmt.Errorf("%w: %v", controllers.ErrDataBindError, err)
		middleware2.AddGinError(ctx, wrappedErr)

		return
	}

	response, err := u.user.SignUp(ctx, req)
	if err != nil {
		middleware2.AddGinError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
