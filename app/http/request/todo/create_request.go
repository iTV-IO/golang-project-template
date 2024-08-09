package todo_request

type TodoCreateRequest struct {
	Title string `json:"title" binding:"required"`
}