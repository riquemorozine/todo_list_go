package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserId"`
}

func NewUser(Name, Email, Password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       uuid.New().String(),
		Name:     Name,
		Email:    Email,
		Password: string(hashedPassword),
	}

	return user, nil
}
