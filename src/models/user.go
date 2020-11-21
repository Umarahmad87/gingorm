package models

import (
	"errors"

	"app/src/store"

	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;type:varchar(100);default:null" json:"email,omitempty"`
	Password string `gorm:"type:varchar(25)" json:"-"`
	Name     string `gorm:"type:varchar(25)" json:"name,omitempty"`
}

type userManager struct{}

//UserManager manager
var UserManager = new(userManager)

// Validate  User
func (user *User) Validate() error {

	if user.Email == "" {
		return errors.New("Email empty")
	}
	if user.Password == "" {
		return errors.New("Password empty")
	}
	return nil
}

// CreateUser insert user in database
func (userModel userManager) CreateUser(user *User) (uint, error) {

	if err := user.Validate(); err != nil {
		return 0, err
	}

	err := store.Db.Create(user).Error
	if err != nil {

		return 0, err
	}

	return user.ID, nil
}

//Login ...
func (userModel userManager) Login(user User) (authUser User, token string, err error) {
	user.Validate()
	// write custom logic here
	return authUser, "token", nil
}
