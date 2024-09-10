package dependencies

import (
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints/Handlers"
	"gorm.io/gorm"
)

type HandleContainer struct {
	CreateTodo entrypoints.Handler
	CreateUser entrypoints.Handler
	LoginUser  entrypoints.Handler
}

func Start(db *gorm.DB, JwtSecret string, JwtExpiresIn int) *HandleContainer {
	createTodoUseCase := usecases.NewCreateTodoUseCase(db)

	createUserUseCase := usecases.NewCreateUserUseCase(db)
	loginUserUseCase := usecases.NewUserLoginUseCase(db, JwtSecret, JwtExpiresIn)

	apiHandlers := HandleContainer{}

	apiHandlers.CreateTodo = &Handlers.CreateTodoHandler{
		UseCase: &createTodoUseCase,
	}

	apiHandlers.CreateUser = &Handlers.CreateUserHandler{
		UseCase: &createUserUseCase,
	}
	apiHandlers.LoginUser = &Handlers.LoginUserHandler{
		UseCase: &loginUserUseCase,
	}

	return &apiHandlers
}
