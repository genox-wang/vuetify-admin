package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
)

// User is user model
type User struct {
	gorm.Model
	DisplayName string `json:"display_name"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `gorm:"not null"`
}

// Login check username and password
func (u *User) Login() error {
	var tmpU User
	if DB.Where("username=?", u.Username).First(&tmpU).Error != nil {
		return errors.New("Username doesn't exist")
	}

	if tmpU.Password == u.CrptoPassword(u.Password) {
		u.DisplayName = tmpU.DisplayName
		u.ID = tmpU.ID
		return nil
	}
	return errors.New("Password is incorrent")
}

// Create create a new user
func (u *User) Create() error {
	u.Password = u.CrptoPassword(u.Password)
	return DB.Create(u).Error
}

// CrptoPassword crpto password
func (u *User) CrptoPassword(password string) string {
	b := sha1.New()
	io.WriteString(b, password)
	return fmt.Sprintf("%x", b.Sum(nil))
}
