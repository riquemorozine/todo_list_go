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

func (t *Todo) FindAll() (*[]entities.Todo, error) {
	todos := &[]entities.Todo{}

	err := t.DB.Find(todos).Error

	if err != nil {
		return nil, err
	}

	// TODO: Pagination and filtering by user id

	return todos, nil
}

func (t *Todo) Update(todo *entities.Todo) error {
	_, err := t.FindByID(todo.ID)

	if err != nil {
		return err
	}

	return t.DB.Save(todo).Error
}

func (t *Todo) Delete(id string) error {
	todo := &entities.Todo{}

	err := t.DB.First(todo, "id = ?", id).Error

	if err != nil {
		return err
	}

	return t.DB.Delete(todo).Error
}
