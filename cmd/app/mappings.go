package app

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/core/middlewares"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureMappings(router *gin.Engine, apiHandlers *dependencies.HandleContainer) *gin.Engine {
	router.Static("/api", "./api")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/swagger.yaml")))

	router.POST("/users", apiHandlers.CreateUser.Handle)
	router.POST("/login", apiHandlers.LoginUser.Handle)

	todos := router.Group("/todos", middlewares.AuthMiddleware())
	todos.POST("/", apiHandlers.CreateTodo.Handle)
	todos.GET("/", apiHandlers.FindAllTodo.Handle)
	todos.PUT("/:id", apiHandlers.UpdateTodo.Handle)
	todos.DELETE("/:id", apiHandlers.DeleteTodo.Handle)

	return router
}
