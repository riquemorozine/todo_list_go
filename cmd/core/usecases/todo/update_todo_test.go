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

func TestUpdateTodo(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not set up test database: %v", err)
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Todo{})
	if err != nil {
		t.Fatalf("could not migrate tables: %v", err)
	}

	todo := entities.NewTodo("Test Title", "Test Description", "pending", "1")
	err = db.Create(todo).Error

	if err != nil {
		t.Fatalf("could not create todo: %v", err)
	}

	usecase := NewUpdateTodoUseCase(db)
	res, err := usecase.Execute(
		context.Background(), &contracts.UpdateTodoRequest{
			Title:       "Test Title Updated",
			Description: "Test Description Updated",
			Status:      "done",
		}, todo.ID)
	if err != nil {
		t.Fatalf("could not update todo: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Test Title Updated", res.Title)
	assert.Equal(t, "Test Description Updated", res.Description)
	assert.Equal(t, "done", res.Status)
}
