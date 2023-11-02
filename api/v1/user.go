package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userRegister.Register(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userLogin.Login(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userUpdate.Update(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func UserUploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	var userUploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUploadAvatar); err != nil {
		c.JSON(http.StatusBadGateway, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userUploadAvatar.UploadAvatar(c.Request.Context(), claims.ID, file, fileHeader)
	c.JSON(http.StatusOK, res)
}

func UserSendingEmail(c *gin.Context) {
	var userSendEmail service.UserSendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userSendEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userSendEmail.UserSendEmail(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func UserValidEmail(c *gin.Context) {
	var userValidEmail service.UserValidEmailService
	if err := c.ShouldBind(&userValidEmail); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return

	}
	res := userValidEmail.UserValidEmail(c.Request.Context(), c.GetHeader("Authorization"))
	c.JSON(http.StatusOK, res)
}

func ShowUserMoney(c *gin.Context) {
	var showUserMoney service.ShowUserMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showUserMoney); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return
	}
	res := showUserMoney.ShowUserMoney(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}
