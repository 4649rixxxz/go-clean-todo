package settings

import (
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

func returnAbortWith(ctx *gin.Context, code int, errors []ErrorResponse) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"errors": errors,
	})
}

func ReturnStatusBadRequest(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusBadRequest, errors)
}

func ReturnStatusUnauthorized(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusUnauthorized, errors)
}

func ReturnStatusForbidden(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusForbidden, errors)
}

func ReturnStatusNotFound(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusNotFound, errors)
}

func ReturnStatusInternalServerError(ctx *gin.Context, errors []ErrorResponse) {
	returnAbortWith(ctx, http.StatusInternalServerError, errors)
}

func ReturnError(ctx *gin.Context, err error) {
	ctx.Error(err)
}
