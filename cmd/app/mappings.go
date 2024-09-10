package app

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/core/middlewares"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
)

func ConfigureMappings(router *gin.Engine, apiHandlers *dependencies.HandleContainer) *gin.Engine {
	router.POST("/users", apiHandlers.CreateUser.Handle)
	router.POST("/login", apiHandlers.LoginUser.Handle)

	todos := router.Group("/todos", middlewares.AuthMiddleware())
	todos.POST("/", apiHandlers.CreateTodo.Handle)
	todos.GET("/", apiHandlers.FindAll.Handle)

	return router
}
