package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/model"
)

// User vo view objective
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		Avatar:    config.Host + config.HttpPort + config.AvatarPath + user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
