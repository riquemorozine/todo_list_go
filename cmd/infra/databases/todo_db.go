package databases

import (
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"gorm.io/gorm"
)

type Todo struct {
	DB *gorm.DB
}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Create(todo *entities.Todo) error {
	return t.DB.Create(todo).Error
}

func (t *Todo) FindByID(id string) (*entities.Todo, error) {
	todo := &entities.Todo{}

	err := t.DB.First(todo, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return todo, nil
}
