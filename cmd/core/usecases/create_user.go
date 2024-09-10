package usecases

import (
	"context"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
)

type ResponseCreateUser struct {
	ID    string `json:"user_id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, req *contracts.CreateUserRequest) (ResponseCreateUser, error)
}

func NewCreateUserUseCase(db *gorm.DB) CreateUserUseCaseImpl {
	return CreateUserUseCaseImpl{
		UserDB: databases.NewUser(db),
	}
}

type CreateUserUseCaseImpl struct {
	UserDB *databases.User
}

func (imp *CreateUserUseCaseImpl) Execute(ctx context.Context, req *contracts.CreateUserRequest) (ResponseCreateUser, error) {
	newUser := entities.NewUser(req.Name, req.Email, req.Password)

	err := imp.UserDB.Create(newUser)

	if err != nil {
		return ResponseCreateUser{}, err
	}

	return ResponseCreateUser{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}
