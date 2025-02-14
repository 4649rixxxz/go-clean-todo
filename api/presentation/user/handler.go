package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-clean-todo/presentation/settings"
	userUsecase "go-clean-todo/usecase/user"
)

type handler struct {
	SignupUserUsecase *userUsecase.SignupUserUsecase
	SigninUserUsecase *userUsecase.SigninUserUsecase
}

func NewHandler(
	SignupUserUsecase *userUsecase.SignupUserUsecase,
	SigninUserUsecase *userUsecase.SigninUserUsecase,
) handler {
	return handler{
		SignupUserUsecase: SignupUserUsecase,
		SigninUserUsecase: SigninUserUsecase,
	}
}

const cookieName = "go-clean-todo"

func (h handler) Signup(ctx *gin.Context) {
	var body SignupParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		settings.ReturnStatusBadRequestForInvalidBody(ctx, err)
		return
	}

	dto := userUsecase.SignupUserUsecaseDTO{
		Email:    body.Email,
		Password: body.Password,
	}
	if err := h.SignupUserUsecase.Run(dto); err != nil {
		settings.ConvertUsecaseErrorToHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h handler) Signin(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := ctx.Bind(&body); err != nil {
		settings.ReturnStatusBadRequestForInvalidBody(ctx, err)
		return
	}
	dto := userUsecase.SigninUserUsecaseDTO{
		Email:    body.Email,
		Password: body.Password,
	}
	jwtToken, err := h.SigninUserUsecase.Run(dto)
	if err != nil {
		settings.ConvertUsecaseErrorToHTTPError(ctx, err)
		return
	}

	ctx.SetCookie(cookieName, jwtToken, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})
}

func (h handler) Signout(ctx *gin.Context) {
	ctx.SetCookie(cookieName, "", 0, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}
