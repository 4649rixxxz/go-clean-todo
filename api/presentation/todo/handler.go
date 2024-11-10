package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-clean-todo/presentation/settings"
	todoUsecase "go-clean-todo/usecase/todo"
)

type handler struct {
	createTodoUsecase *todoUsecase.CreateTodoUsecase
}

func NewHandler(
	createTodoUsecase *todoUsecase.CreateTodoUsecase,
) handler {
	return handler{
		createTodoUsecase: createTodoUsecase,
	}
}

func (h handler) Create(ctx *gin.Context) {
	// データ型のバリデーション
	var body CreateTodoParams
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	userIDFromCtx, isExisting := ctx.Get("user_id")
	// Todo 書き方が冗長なので修正したい
	if !isExisting {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create todo",
		})

		return
	}
	userID, ok := userIDFromCtx.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create todo",
		})

		return
	}
	inputDTO := todoUsecase.CreateTodoUsecaseInputDTO{
		UserID:      userID,
		Title:       body.Title,
		Description: body.Description,
	}
	outputDTO, err := h.createTodoUsecase.Run(inputDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create todo",
		})

		return
	}

	response := createTodoResponse{
		Todo: createTodoResponseModel{
			TodoID:           outputDTO.TodoID,
			UserID:           outputDTO.UserID,
			Title:            outputDTO.Title,
			Description:      outputDTO.Description,
			AttachedFilePath: outputDTO.AttachedFilePath,
			CompletedAt:      outputDTO.CompletedAt,
			CreatedAt:        outputDTO.CreatedAt,
			UpdatedAt:        outputDTO.UpdatedAt,
		},
	}

	settings.ReturnStatusCreated(ctx, response)
}
