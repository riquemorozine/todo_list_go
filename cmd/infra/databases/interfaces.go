package databases

import "github.com/riquemorozine/todo_list_go/cmd/entities"

type TodoInterface interface {
	Create(todo *entities.Todo) error
	GetAll() (*[]entities.Todo, error)
	FindByID(id string) (*entities.Todo, error)
	FindByName(name string) (*[]entities.Todo, error)
	Delete(id string) error
}

type UserInterface interface {
	Create(user *entities.User) error
	FindById(id string) (*entities.User, error)
	Delete(id string) error
}
