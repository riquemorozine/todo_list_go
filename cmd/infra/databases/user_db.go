package databases

import (
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByID(id string) (*entities.User, error) {
	user := &entities.User{}

	err := u.DB.First(user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Delete(id string) error {
	user := &entities.User{}

	err := u.DB.First(user, "id = ?", id).Error

	if err != nil {
		return err
	}

	return u.DB.Delete(user).Error
}
