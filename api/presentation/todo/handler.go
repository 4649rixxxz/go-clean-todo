package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
