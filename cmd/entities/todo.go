package entities

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserId      string
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(Title, Description, Status, UserId string) *Todo {
	todo := &Todo{
		ID:          uuid.New().String(),
		Title:       Title,
		Description: Description,
		Status:      Status,
		UserId:      UserId,
		CreatedAt:   time.Now(),
	}

	return todo
}
