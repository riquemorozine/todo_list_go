package dependencies

import (
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints/Handlers"
	"gorm.io/gorm"
)

type HandleContainer struct {
	CreateTodo entrypoints.Handler
}

func Start(db *gorm.DB) *HandleContainer {

	createTodoUseCase := usecases.NewCreateTodoUseCase(db)

	apiHandlers := HandleContainer{}

	apiHandlers.CreateTodo = &Handlers.CreateTodoHandler{
		UseCase: &createTodoUseCase,
	}

	return &apiHandlers
}
