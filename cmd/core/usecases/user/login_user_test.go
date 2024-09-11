package user

import (
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestLoginUser(t *testing.T) {
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

	usecase := NewUserLoginUseCase(db, "secret", 3000)
	res, err := usecase.Execute(&contracts.LoginUserRequest{
		Email:    user.Email,
		Password: "1234",
	})
	if err != nil {
		t.Fatalf("could not login user: %v", err)
	}

	if res.AccessToken == "" {
		t.Fatalf("could not login user: token is empty")
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, res.AccessToken)
}
