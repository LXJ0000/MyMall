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
		c.JSON(http.StatusBadRequest, err)
	}
	res := userRegister.Register(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	res := userLogin.Login(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	res := userUpdate.Update(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func UserUploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	var userUploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUploadAvatar); err != nil {
		c.JSON(http.StatusBadGateway, err)
	}
	res := userUploadAvatar.UploadAvatar(c.Request.Context(), claims.ID, file, fileHeader)
	c.JSON(http.StatusOK, res)
}

func UserSendingEmail(c *gin.Context) {
	
}
