package dependencies

import (
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints/Handlers"
)

type HandleContainer struct {
	CreateTodo entrypoints.Handler
}

func Start() *HandleContainer {
	createTodoUseCase := usecases.NewCreateTodoUseCase()

	apiHandlers := HandleContainer{}

	apiHandlers.CreateTodo = &Handlers.CreateTodoHandler{
		UseCase: &createTodoUseCase,
	}

	return &apiHandlers
}
