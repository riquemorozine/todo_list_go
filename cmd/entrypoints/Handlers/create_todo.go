package Handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases"
	"github.com/riquemorozine/todo_list_go/cmd/errors"
	"net/http"
)

type CreateTodoHandler struct {
	UseCase usecases.CreateTodoUseCase
}

func (handler *CreateTodoHandler) Handle(c *gin.Context) {
	err := handler.handle(c)

	if err != nil {
		c.JSON(err.Status, err)
	}

	return
}

func (handler *CreateTodoHandler) handle(c *gin.Context) *errors.APIError {
	ctx := c.Request.Context()

	r, err := handler.UseCase.Execute(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)

	return nil
}
