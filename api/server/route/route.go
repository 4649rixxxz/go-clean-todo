package route

import (
	ginpkg "github.com/gin-gonic/gin"

	"go-clean-todo/infrastructure/mysql/repository"
	"go-clean-todo/presentation/middleware"
	todoHandler "go-clean-todo/presentation/todo"
	userHandler "go-clean-todo/presentation/user"
	todoUsecase "go-clean-todo/usecase/todo"
	userUsecace "go-clean-todo/usecase/user"
)

func InitRoute(api *ginpkg.Engine) {
	v1 := api.Group("/v1")
	{
		userRoute(v1)
		todoRoute(v1)
	}
}

func userRoute(r *ginpkg.RouterGroup) {
	userRepo := repository.NewUserRepository()
	handler := userHandler.NewHandler(
		userUsecace.NewSignupUserUsecase(userRepo),
		userUsecace.NewSigninUserUsecase(userRepo),
	)
	r.POST("/signup", handler.Signup)
	r.POST("/signin", handler.Signin)
	r.POST("/signout", handler.Signout)
}

func todoRoute(r *ginpkg.RouterGroup) {
	todos := r.Group("/todos", middleware.Auth())
	todoRepo := repository.NewTodoRepository()
	handler := todoHandler.NewHandler(
		todoUsecase.NewCreateTodoUsecase(todoRepo),
	)
	todos.POST("", handler.Create)
}
