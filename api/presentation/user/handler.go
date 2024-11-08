package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

func (h handler) Signup(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	dto := userUsecase.SignupUserUsecaseDTO{
		Email:    body.Email,
		Password: body.Password,
	}
	if err := h.SignupUserUsecase.Run(dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h handler) Signin(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	dto := userUsecase.SigninUserUsecaseDTO{
		Email:    body.Email,
		Password: body.Password,
	}
	jwtToken, err := h.SigninUserUsecase.Run(dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	ctx.SetCookie("Authorization", jwtToken, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})

}

func (h handler) Signout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", 0, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}
