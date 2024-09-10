package todo

import (
	"context"
	"fmt"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type FindAllTodoUseCase interface {
	Execute(ctx context.Context, req *contracts.FindAllTodoRequest, userId string) ([]ResponseFindAllTodo, error)
}

type ResponseFindAllTodo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type FindAllTodoUseCaseImpl struct {
	TodoDB *databases.Todo
}

func NewFindAllTodoUseCase(db *gorm.DB) FindAllTodoUseCaseImpl {
	return FindAllTodoUseCaseImpl{
		TodoDB: databases.NewTodo(db),
	}
}

func (imp *FindAllTodoUseCaseImpl) Execute(ctx context.Context, req *contracts.FindAllTodoRequest, userId string) ([]ResponseFindAllTodo, error) {

	fmt.Println(req.PageSize, req.Page, req.Sort)
	todos, err := imp.TodoDB.FindAll(userId, req.Sort, req.Page, req.PageSize)

	if err != nil {
		return []ResponseFindAllTodo{}, err
	}

	var res []ResponseFindAllTodo

	for _, todo := range *todos {
		res = append(res, ResponseFindAllTodo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Status:      todo.Status,
			CreatedAt:   todo.CreatedAt.String(),
		})
	}

	return res, nil
}
