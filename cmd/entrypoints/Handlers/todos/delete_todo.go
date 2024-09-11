package todos

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases/todo"
	"github.com/riquemorozine/todo_list_go/cmd/errors"
	"net/http"
)

type DeleteTodoHandler struct {
	UseCase todo.DeleteTodoUseCase
}

func (handler *DeleteTodoHandler) Handle(c *gin.Context) {
	err := handler.handle(c)

	if err != nil {
		c.JSON(err.Status, err)
	}

	return
}

func (handler *DeleteTodoHandler) handle(c *gin.Context) *errors.APIError {
	ctx := c.Request.Context()
	id := c.Param("id")

	err := handler.UseCase.Execute(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
	return nil
}
