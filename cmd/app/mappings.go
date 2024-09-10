package app

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/core/middlewares"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
)

func ConfigureMappings(router *gin.Engine, apiHandlers *dependencies.HandleContainer) *gin.Engine {
	router.POST("/users", apiHandlers.CreateUser.Handle)
	router.POST("/login", apiHandlers.LoginUser.Handle)

	router.Use(middlewares.AuthMiddleware())
	router.POST("/todos", apiHandlers.CreateTodo.Handle)

	return router
}
