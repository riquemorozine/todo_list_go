package todo

import (
	"context"
	"fmt"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestDeleteTodo(t *testing.T) {
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

	usecase := NewDeleteTodoUseCase(db)
	err = usecase.Execute(context.Background(), todo.ID)
	if err != nil {
		t.Fatalf("could not delete todo: %v", err)
	}

	_, err = usecase.TodoDB.FindByID(todo.ID)
	fmt.Println(err)

	if err == nil {
		t.Fatalf("todo was not deleted")
	}

	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
}
