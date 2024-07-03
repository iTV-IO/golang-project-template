package request

type TodoCreate struct {
	Title string `json:"title" binding:"required"`
}
