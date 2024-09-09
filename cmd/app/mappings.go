package app

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
)

func ConfigureMappings(router *gin.Engine, apiHandlers *dependencies.HandleContainer) *gin.Engine {
	router.POST("/todos", apiHandlers.CreateTodo.Handle)

	return router
}
