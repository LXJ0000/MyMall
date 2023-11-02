package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user模块错误
	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorPassword              = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeOut = 30006
	ErrorFileUploadFail        = 30007
	ErrorSendMailFail          = 30008

	// product模块错误
	ErrorProductImgUpload = 40001
)
