package todo

import "time"

type createTodoResponseModel struct {
	TodoID           uint       `json:"todo_id"`
	UserID           uint       `json:"user_id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	AttachedFilePath *string    `json:"attached_file_path"`
	CompletedAt      *time.Time `json:"completed_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type createTodoResponse struct {
	Todo createTodoResponseModel `json:"todo"`
}
