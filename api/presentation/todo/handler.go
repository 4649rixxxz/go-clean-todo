package todo

import (
	"github.com/gin-gonic/gin"

	"go-clean-todo/pkg/validator"
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
	var body CreateTodoParams
	if err := ctx.ShouldBindJSON(&body); err != nil {
		settings.ReturnStatusBadRequestForInvalidBody(ctx, err)
		return
	}
	validate := validator.GetValidator()
	if err := validate.Struct(body); err != nil {
		errMsgs := validator.MakeValidationErrMessages(err)
		settings.ReturnStatusBadRequest(ctx, errMsgs)
		return
	}
	userIDFromCtx, isExisting := ctx.Get("user_id")
	// Todo 書き方が冗長なので修正したい
	if !isExisting {
		settings.ReturnStatusInternalServerError(ctx, "Failed to create todo")
		return
	}
	userID, ok := userIDFromCtx.(uint)
	if !ok {
		settings.ReturnStatusInternalServerError(ctx, "Failed to create todo")
		return
	}
	inputDTO := todoUsecase.CreateTodoUsecaseInputDTO{
		UserID:      userID,
		Title:       body.Title,
		Description: body.Description,
	}
	outputDTO, err := h.createTodoUsecase.Run(inputDTO)
	if err != nil {
		settings.ConvertUsecaseErrorToHTTPError(ctx, err)
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
