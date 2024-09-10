package dependencies

import (
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases/todo"
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases/user"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints"
	todo2 "github.com/riquemorozine/todo_list_go/cmd/entrypoints/Handlers/todos"
	"github.com/riquemorozine/todo_list_go/cmd/entrypoints/Handlers/users"
	"gorm.io/gorm"
)

type HandleContainer struct {
	CreateTodo entrypoints.Handler
	FindAll    entrypoints.Handler
	CreateUser entrypoints.Handler
	LoginUser  entrypoints.Handler
}

func Start(db *gorm.DB, JwtSecret string, JwtExpiresIn int) *HandleContainer {
	createTodoUseCase := todo.NewCreateTodoUseCase(db)
	findAllTodoUseCase := todo.NewFindAllTodoUseCase(db)

	createUserUseCase := user.NewCreateUserUseCase(db)
	loginUserUseCase := user.NewUserLoginUseCase(db, JwtSecret, JwtExpiresIn)

	apiHandlers := HandleContainer{}

	apiHandlers.CreateTodo = &todo2.CreateTodoHandler{
		UseCase: &createTodoUseCase,
	}
	apiHandlers.FindAll = &todo2.FindAllTodosHandler{
		UseCase: &findAllTodoUseCase,
	}

	apiHandlers.CreateUser = &users.CreateUserHandler{
		UseCase: &createUserUseCase,
	}
	apiHandlers.LoginUser = &users.LoginUserHandler{
		UseCase: &loginUserUseCase,
	}

	return &apiHandlers
}
