package v1

import (
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
