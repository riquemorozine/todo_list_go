package usecases

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type ResponseCreateTodo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type CreateTodoUseCase interface {
	Execute(ctx context.Context, req *contracts.CreateTodoRequest) (ResponseCreateTodo, error)
}

func NewCreateTodoUseCase(db *gorm.DB) CreateTodoUseCaseImpl {
	return CreateTodoUseCaseImpl{
		TodoDB: databases.NewTodo(db),
	}
}

type CreateTodoUseCaseImpl struct {
	TodoDB *databases.Todo
}

func (imp *CreateTodoUseCaseImpl) Execute(ctx context.Context, req *contracts.CreateTodoRequest) (ResponseCreateTodo, error) {

	td := entities.NewTodo(req.Title, req.Description, req.Status, "124")

	err := imp.TodoDB.Create(td)

	if err != nil {
		return ResponseCreateTodo{}, err
	}

	return ResponseCreateTodo{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Status:      td.Status,
		CreatedAt:   td.CreatedAt.String(),
	}, nil
}
