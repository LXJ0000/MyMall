package e

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "参数错误",
	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorPassword:              "密码错误",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeOut: "token过期",
	ErrorFileUploadFail:        "文件上传失败",
}

// GetMsg 获取状态码信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
