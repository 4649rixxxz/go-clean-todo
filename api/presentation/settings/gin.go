package settings

import (
	"go-clean-todo/usecase"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router.Use(cors.New(config))
	return router
}

func ReturnStatusOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusCreated, &body)
}

func ReturnStatusNoContent(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func ReturnStatusBadRequestForInvalidBody(ctx *gin.Context, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func returnAbortWith(ctx *gin.Context, code int, message string, errors []ErrorResponse) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"errors":  errors,
	})
}

func ReturnStatusBadRequest(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusBadRequest, "Bad Request", errors)
}

func ReturnStatusUnauthorized(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusUnauthorized, "Unauthorized", errors)
}

func ReturnStatusForbidden(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusForbidden, "Forbidden", errors)
}

func ReturnStatusNotFound(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusNotFound, "Not Found", errors)
}

func ReturnStatusInternalServerError(ctx *gin.Context, message string) {
	returnAbortWith(ctx, http.StatusInternalServerError, message, []ErrorResponse{})
}

func ReturnError(ctx *gin.Context, err error) {
	ctx.Error(err)
}

func ConvertUsecaseErrorToHTTPError(ctx *gin.Context, e usecase.UsecaseErrorI) {
	switch e.Code() {
	case usecase.InvalidInputError:
		ReturnStatusBadRequest(
			ctx,
			[]ErrorResponse{
				{Field: e.Field(), Message: e.Error()},
			},
		)
	case usecase.ResourceNotFoundError:
		ReturnStatusNotFound(
			ctx,
			[]ErrorResponse{
				{Field: e.Field(), Message: e.Error()},
			},
		)
	default:
		ReturnStatusInternalServerError(
			ctx,
			e.Error(),
		)
	}
}
