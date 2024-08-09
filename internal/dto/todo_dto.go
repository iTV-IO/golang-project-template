package dto

type TodoCreateDTO struct {
	Title string `json:"title" binding:"required"`
}

type TodoOutputDTO struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
