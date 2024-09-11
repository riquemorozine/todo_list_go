package todo

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type DeleteTodoUseCase interface {
	Execute(ctx context.Context, todoId string) error
}

func NewDeleteTodoUseCase(db *gorm.DB) DeleteTodoUseCaseImpl {
	return DeleteTodoUseCaseImpl{
		TodoDB: *databases.NewTodo(db),
	}
}

type DeleteTodoUseCaseImpl struct {
	TodoDB databases.Todo
}

func (imp *DeleteTodoUseCaseImpl) Execute(ctx context.Context, todoId string) error {
	todo, err := imp.TodoDB.FindByID(todoId)
	if err != nil {
		return err
	}

	err = imp.TodoDB.Delete(todo.ID)
	if err != nil {
		return err
	}

	return nil
}
