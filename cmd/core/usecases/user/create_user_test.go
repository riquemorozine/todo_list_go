package user

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("could not set up test database: %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Todo{})

	req := &contracts.CreateUserRequest{
		Name:     "Test Name",
		Email:    "test@email.com",
		Password: "Test Password",
	}

	usecase := NewCreateUserUseCase(db)

	response, err := usecase.Execute(context.Background(), req)

	if err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	res, err := usecase.UserDB.FindByID(response.ID)
	if err != nil {
		t.Fatalf("could not find user: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Email, res.Email)
	assert.NotEqual(t, req.Password, res.Password)
}
