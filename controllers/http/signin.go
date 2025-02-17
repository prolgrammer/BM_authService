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

type signInController struct {
	user usecases.SignInUseCase
}

func NewSignInController(
	handler *gin.Engine,
	signInUseCase usecases.SignInUseCase,
	middleware middleware2.Middleware,
) {

	u := &signInController{
		user: signInUseCase,
	}

	handler.POST("/signin", u.SignIn, middleware.HandleErrors)
}

// SignIn godoc
// @Summary вход в аккаунт
// @Description вход в аккаунт по почте + паролю
// @Accept json
// @Produce json
// @Param request body requests.SignRequest true "структура запроса"
// @Success 200 {object} responses.SignResponse
// @Failure 400 {object} string "некорректный формат запроса"
// @Failure 401 {object} string "неправильный пароль"
// @Failure 404 {object} string "пользователь не найден"
// @Failure 500 {object} string "внутренняя ошибка сервера"
// @Router /signin [post]
func (uc *signInController) SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var request requests.SignRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		wrappedError := fmt.Errorf("%w: %w", controllers.ErrDataBindError, err)
		middleware2.AddGinError(ctx, wrappedError)
		return
	}

	response, err := uc.user.SignIn(ctx, request)
	if err != nil {
		middleware2.AddGinError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
