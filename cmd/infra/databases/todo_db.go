package databases

import (
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"gorm.io/gorm"
)

type Todo struct {
	DB *gorm.DB
}

func NewTodo(db *gorm.DB) *Todo {
	return &Todo{DB: db}
}

func (t *Todo) Create(todo *entities.Todo) (*entities.Todo, error) {
	err := t.DB.Create(todo).Error

	return todo, err
}

func (t *Todo) FindByID(id string) (*entities.Todo, error) {
	todo := &entities.Todo{}

	err := t.DB.First(todo, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return todo, nil
}
