package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prolgrammer/BM_authService/controllers/requests"
	"github.com/prolgrammer/BM_authService/internal/usecases"
	e "github.com/prolgrammer/BM_package/errors"
	"github.com/prolgrammer/BM_package/middleware"
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

	handler.POST("/auth/signin", u.SignIn, middleware.HandleErrors)
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
// @Router /auth/signin [post]
func (uc *signInController) SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var request requests.SignRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		wrappedError := fmt.Errorf("%w: %w", e.ErrDataBindError, err)
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
