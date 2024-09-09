package usecases

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type CreateTodoUseCase interface {
	Execute(ctx context.Context, req *contracts.CreateTodoRequest) (interface{}, error)
}

func NewCreateTodoUseCase(db *gorm.DB) CreateTodoUseCaseImpl {
	return CreateTodoUseCaseImpl{
		TodoDB: databases.NewTodo(db),
	}
}

type CreateTodoUseCaseImpl struct {
	TodoDB *databases.Todo
}

func (imp *CreateTodoUseCaseImpl) Execute(ctx context.Context, req *contracts.CreateTodoRequest) (interface{}, error) {
	td := entities.NewTodo(req.Title, req.Description, req.Status)

	if err := imp.TodoDB.Create(td); err != nil {
		return nil, err
	}

	return req, nil
}
