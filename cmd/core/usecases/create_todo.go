package usecases

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type CreateTodoUseCase interface {
	Execute(ctx context.Context) (interface{}, error)
}

func NewCreateTodoUseCase(db *gorm.DB) CreateTodoUseCaseImpl {
	return CreateTodoUseCaseImpl{
		TodoDB: databases.NewTodo(),
	}
}

type CreateTodoUseCaseImpl struct {
	TodoDB *databases.Todo
}

func (imp *CreateTodoUseCaseImpl) Execute(ctx context.Context) (interface{}, error) {

	defer ctx.Done()

	return nil, nil
}
