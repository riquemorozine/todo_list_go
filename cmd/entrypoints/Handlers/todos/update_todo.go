package todos

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/config"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases/todo"
	"github.com/riquemorozine/todo_list_go/cmd/errors"
	"net/http"
)

type UpdateTodoHandler struct {
	UseCase todo.UpdateTodoUseCase
}

func (handler *UpdateTodoHandler) Handle(c *gin.Context) {
	err := handler.handle(c)

	if err != nil {
		c.JSON(err.Status, err)
	}

	return
}

func (handler *UpdateTodoHandler) handle(c *gin.Context) *errors.APIError {
	ctx := c.Request.Context()
	id := c.Param("id")

	req := contracts.UpdateTodoRequest{}
	causes, err := config.BindAndValidate(c, &req)

	if err != nil {
		return errors.NewBadRequestError("some fields are invalid", causes)
	}

	r, err := handler.UseCase.Execute(ctx, &req, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)
	return nil
}
