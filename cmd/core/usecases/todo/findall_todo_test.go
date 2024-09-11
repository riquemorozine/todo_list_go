package todo

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestFindAllTodo(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not set up test database: %v", err)
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Todo{})
	if err != nil {
		t.Fatalf("could not migrate tables: %v", err)
	}

	user, err := entities.NewUser("test", "test@email.com", "1234")
	if err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	err = db.Create(user).Error
	if err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	todo1 := entities.NewTodo("Test Title 1", "Test Description 1", "pending", user.ID)
	todo2 := entities.NewTodo("Test Title 2", "Test Description 2", "pending", user.ID)
	todo3 := entities.NewTodo("Test Title 3", "Test Description 3", "pending", user.ID)
	todos := []*entities.Todo{
		todo1,
		todo2,
		todo3,
	}
	for _, todo := range todos {
		err = db.Create(todo).Error
		if err != nil {
			t.Fatalf("could not create todo: %v", err)
		}
	}

	usecase := NewFindAllTodoUseCase(db)
	res, err := usecase.Execute(context.Background(), &contracts.FindAllTodoRequest{PageSize: 1, Page: 1, Sort: "asc"}, user.ID)
	if err != nil {
		t.Fatalf("could not find todos: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, todo1.Title, res[0].Title)
	assert.Equal(t, todo1.Description, res[0].Description)
	assert.Equal(t, todo1.Status, res[0].Status)
}
