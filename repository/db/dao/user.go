package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDbClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 判断用户名是否存在
func (u *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = u.DB.Model(&model.User{}).Where("username=?", userName).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

func (u *UserDao) CreateUser(user *model.User) error {
	return u.DB.Model(&model.User{}).Create(&user).Error
}
