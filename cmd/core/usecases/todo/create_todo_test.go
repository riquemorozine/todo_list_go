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

func TestCreateTodoUseCase_Execute(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("could not set up test database: %v", err)
	}

	db.AutoMigrate(&entities.Todo{})

	req := &contracts.CreateTodoRequest{
		Title:       "Test Title",
		Description: "Test Description",
		Status:      "pending",
	}

	usecase := NewCreateTodoUseCase(db)

	response, err := usecase.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("could not create todos: %v", err)
	}

	res, err := usecase.TodoDB.FindByID(response.ID)
	if err != nil {
		t.Fatalf("could not find todos: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
	assert.Equal(t, req.Status, res.Status)
}
