package todo

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type UpdateTodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTodoUseCase interface {
	Execute(ctx context.Context, req *contracts.UpdateTodoRequest, todoId string) (*UpdateTodoResponse, error)
}

func NewUpdateTodoUseCase(db *gorm.DB) UpdateTodoUseCaseImpl {
	return UpdateTodoUseCaseImpl{
		TodoDB: databases.NewTodo(db),
	}
}

type UpdateTodoUseCaseImpl struct {
	TodoDB *databases.Todo
}

func (imp *UpdateTodoUseCaseImpl) Execute(ctx context.Context, req *contracts.UpdateTodoRequest, todoId string) (*UpdateTodoResponse, error) {
	todo, err := imp.TodoDB.FindByID(todoId)

	if err != nil {
		return nil, err
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Status = req.Status
	err = imp.TodoDB.Update(todo)

	if err != nil {
		return nil, err
	}

	return &UpdateTodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
	}, nil
}
