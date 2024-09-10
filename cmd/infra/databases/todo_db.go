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

func (t *Todo) FindAll(userId string, sort string, page, pageSize int) (*[]entities.Todo, error) {
	todos := &[]entities.Todo{}
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && pageSize != 0 {
		err = t.DB.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at "+sort).Where("user_id = ?", userId).Find(todos).Error
	} else {
		err = t.DB.Where("user_id = ?", userId).Find(todos).Error
	}

	return todos, err
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
