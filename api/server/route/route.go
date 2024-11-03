package route

import (
	"go-clean-todo/infrastructure/mysql/repository"
	userHandler "go-clean-todo/presentation/user"
	userUsecace "go-clean-todo/usecase/user"

	ginpkg "github.com/gin-gonic/gin"
)

func InitRoute(api *ginpkg.Engine) {
	v1 := api.Group("/v1")
	{
		userRoute(v1)
	}
}

func userRoute(r *ginpkg.RouterGroup) {
	userRepo := repository.NewUserRepository()
	h := userHandler.NewHandler(
		userUsecace.NewSignupUserUsecase(userRepo),
		userUsecace.NewSigninUserUsecase(userRepo),
		userUsecace.NewSignoutUserUsecase(userRepo),
	)
	r.POST("/signup", h.Signup)
	r.POST("/signin", h.Signin)
	r.POST("/signout", h.Signout)
}
