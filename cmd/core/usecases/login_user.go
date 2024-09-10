package usecases

import (
	"github.com/go-chi/jwtauth"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gorm.io/gorm"
	"time"
)

type ResponseUserLogin struct {
	AccessToken string `json:"access_token"`
}

type UserLoginUseCase interface {
	Execute(req *contracts.LoginUserRequest) (ResponseUserLogin, error)
}

func NewUserLoginUseCase(db *gorm.DB, jwt *jwtauth.JWTAuth, jwtExpiresIn int) UserLoginUseCaseImpl {
	return UserLoginUseCaseImpl{
		DB:           databases.NewUser(db),
		JWT:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

type UserLoginUseCaseImpl struct {
	DB           *databases.User
	JWT          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func (imp *UserLoginUseCaseImpl) Execute(req *contracts.LoginUserRequest) (ResponseUserLogin, error) {
	user, err := imp.DB.FindByEmail(req.Email)

	if err != nil {
		return ResponseUserLogin{}, err
	}

	if !user.ComparePassword(req.Password) {
		return ResponseUserLogin{}, err
	}

	_, token, _ := imp.JWT.Encode(map[string]interface{}{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(imp.JwtExpiresIn)).Unix(),
	})

	return ResponseUserLogin{
		AccessToken: token,
	}, nil
}
