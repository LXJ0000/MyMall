package service

import (
	"MyMall/pkg/e"
	util "MyMall/pkg/utils"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

func (u *UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	if u.Key == "" || len(u.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	//	密文存储金额
	util.Encrypt.SetKey(u.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, _ := userDao.ExistOrNotByUserName(u.UserName)
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		UserName: u.UserName,
		NickName: u.NickName,
		Avatar:   "avatar.jpg",
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额 加密

	}
	if err := user.SetPassword(u.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// Create User
	if err := userDao.CreateUser(user); err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (u *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	//判断用户是否存在
	user, exist, _ := userDao.ExistOrNotByUserName(u.UserName)
	if !exist {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册。",
		}
	}
	//校验密码
	if !user.CheckPassword(u.Password) {
		code = e.ErrorPassword
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重试。",
		}
	}

	//http 无状态（认证：token）
	token, err := util.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重试。",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}
