package usecases

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func NewTodoDBMock(db *gorm.DB) *databases.Todo {
	return &databases.Todo{
		DB: db,
	}
}

func TestCreateTodoUseCase_Execute(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("could not set up test database: %v", err)
	}

	db.AutoMigrate(&entities.Todo{})

	useCase := CreateTodoUseCaseImpl{
		TodoDB: NewTodoDBMock(db),
	}

	req := &contracts.CreateTodoRequest{
		Title:       "Test Title",
		Description: "Test Description",
		Status:      "pending",
	}

	response, err := useCase.Execute(context.Background(), req)

	var saveTodo entities.Todo
	if err := db.First(&saveTodo, "id = ?", response.ID).Error; err != nil {
		t.Fatalf("could not find todo: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, req.Title, response.Title)
	assert.Equal(t, req.Description, response.Description)
	assert.Equal(t, req.Status, response.Status)
}
