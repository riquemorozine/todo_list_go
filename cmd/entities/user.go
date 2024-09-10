package entities

import "github.com/google/uuid"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserID"`
}

func NewUser(Name, Email, Password string) *User {
	user := &User{
		ID:       uuid.New().String(),
		Name:     Name,
		Email:    Email,
		Password: Password,
	}

	return user
}
