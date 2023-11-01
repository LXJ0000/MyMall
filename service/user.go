package service

import (
	"MyMall/config"
	"MyMall/pkg/e"
	util "MyMall/pkg/utils"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"fmt"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"time"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

type UserSendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	//1 绑定邮箱 2解绑邮箱 3修改密码
}

type UserValidEmailService struct {
}

type ShowUserMoneyService struct {
	Key string `json:"key" form:"key"`
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
			Error:  "用户不存在",
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
			Error:  "密码解析失败",
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

func (u *UserService) Update(ctx context.Context, userId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	//获取user
	user, _ := userDao.GetUserByUserId(userId)
	//	修改nick_name
	if u.NickName != "" {
		user.NickName = u.NickName
	}

	if err := userDao.UpdateUserByUserId(userId, user); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User: serializer.BuildUser(user),
		},
	}

}

func (u *UserService) UploadAvatar(ctx context.Context, userId uint, file multipart.File, fileHeader *multipart.FileHeader) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, _ := userDao.GetUserByUserId(userId)

	filePath, err := util.UploadAvatarToLocalStatic(user.ID, user.UserName, file, fileHeader.Filename)
	if err != nil {
		code = e.ErrorFileUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   err.Error(),
		}
	}
	user.Avatar = filePath
	if err = userDao.UpdateUserByUserId(userId, user); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
			Error:  "用户信息更新失败-头像",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (u *UserSendEmailService) UserSendEmail(ctx context.Context, userId uint) serializer.Response {
	code := e.Success
	fmt.Println(u)
	token, err := util.GenerateEmailToken(userId, u.Email, u.Password, u.OperationType)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
			Error:  "用户登陆信息过期",
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err := noticeDao.GetNoticeById(u.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
			Error:  "Operation选择错误",
		}
	}
	addr := config.ValidEmail + token
	mailStr := notice.Text
	mailText := mailStr + addr
	m := mail.NewMessage()
	m.SetHeader("From", config.SmtpEmail)
	m.SetHeader("To", u.Email)
	m.SetHeader("Subject", "From LXJ ")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	fmt.Println(m)
	fmt.Println(d)

	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendMailFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (u *UserValidEmailService) UserValidEmail(ctx context.Context, token string) serializer.Response {
	code := e.Success
	var (
		userId        uint
		password      string
		operationType uint
		email         string
	)
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = e.ErrorAuthCheckTokenTimeOut
		} else {
			userId = claims.ID
			password = claims.Password
			operationType = claims.OperationType
			email = claims.Email
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//用户查询
	userDao := dao.NewUserDao(ctx)
	user, _ := userDao.GetUserByUserId(userId)

	if operationType == 1 {
		//绑定
		user.Email = email
	} else if operationType == 2 {
		//解绑
		user.Email = ""
	} else {
		//修改密码
		if err := user.SetPassword(password); err != nil {
			code = e.ErrorPassword
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	_ = userDao.UpdateUserByUserId(userId, user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (s *ShowUserMoneyService) ShowUserMoney(ctx context.Context, userId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByUserId(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户id获取用户信息失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUserMoney(user, s.Key),
		Msg:    e.GetMsg(code),
	}
}
