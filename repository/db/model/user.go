package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email    string
	Password string
	NickName string
	Status   string
	Avatar   string
	Money    string
}

const (
	PasswordCost = 12       // 密码加密难度
	Active       = "active" // 激活用户
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	//hash_password password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true

}
