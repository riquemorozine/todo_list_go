package usecases

import (
	"context"
)

type CreateTodoUseCase interface {
	Execute(ctx context.Context) (interface{}, error)
}

func NewCreateTodoUseCase() CreateTodoUseCaseImpl {
	return CreateTodoUseCaseImpl{}
}

type CreateTodoUseCaseImpl struct {
}

func (imp *CreateTodoUseCaseImpl) Execute(ctx context.Context) (interface{}, error) {
	return nil, nil
}
