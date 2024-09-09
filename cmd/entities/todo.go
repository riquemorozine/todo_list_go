package entities

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(Title, Description, Status string) *Todo {
	todo := &Todo{
		ID:          uuid.New().String(),
		Title:       Title,
		Description: Description,
		Status:      Status,
		CreatedAt:   time.Now(),
	}

	return todo
}
