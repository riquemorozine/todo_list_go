package users

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/config"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/core/usecases/user"
	"github.com/riquemorozine/todo_list_go/cmd/errors"
	"net/http"
)

type CreateUserHandler struct {
	UseCase user.CreateUserUseCase
}

func (handler *CreateUserHandler) Handle(c *gin.Context) {
	err := handler.handle(c)

	if err != nil {
		c.JSON(err.Status, err)
	}
}

func (handler *CreateUserHandler) handle(c *gin.Context) *errors.APIError {
	ctx := c.Request.Context()

	req := contracts.CreateUserRequest{}
	causes, err := config.BindAndValidate(c, &req)

	if err != nil {
		return errors.NewBadRequestError("some fields are invalid", causes)
	}

	r, err := handler.UseCase.Execute(ctx, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)

	return nil
}
