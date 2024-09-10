package user

import (
	"fmt"
	"github.com/riquemorozine/todo_list_go/cmd/core/contracts"
	"github.com/riquemorozine/todo_list_go/cmd/infra/databases"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"gorm.io/gorm"
	"time"
)

type ResponseUserLogin struct {
	AccessToken string `json:"access_token"`
}

type UserLoginUseCase interface {
	Execute(req *contracts.LoginUserRequest) (ResponseUserLogin, error)
}

func NewUserLoginUseCase(db *gorm.DB, JwtSecret string, JwtExpiresIn int) UserLoginUseCaseImpl {
	return UserLoginUseCaseImpl{
		DB:           databases.NewUser(db),
		JwtSecret:    JwtSecret,
		JwtExpiresIn: JwtExpiresIn,
	}
}

type UserLoginUseCaseImpl struct {
	DB           *databases.User
	JwtExpiresIn int
	JwtSecret    string
}

func (imp *UserLoginUseCaseImpl) Execute(req *contracts.LoginUserRequest) (ResponseUserLogin, error) {
	user, err := imp.DB.FindByEmail(req.Email)

	if err != nil {
		return ResponseUserLogin{}, err
	}

	if !user.ComparePassword(req.Password) {
		return ResponseUserLogin{}, err
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(imp.JwtExpiresIn)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(imp.JwtSecret))

	if err != nil {
		return ResponseUserLogin{}, err
	}

	fmt.Println(imp.JwtSecret)

	return ResponseUserLogin{
		AccessToken: tokenString,
	}, nil
}
