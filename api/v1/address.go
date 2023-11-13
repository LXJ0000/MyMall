package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createAddress service.AddressService
	if err := c.ShouldBind(&createAddress); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := createAddress.CreateAddress(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func GetAddress(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getAddress service.AddressService
	if err := c.ShouldBind(&getAddress); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getAddress.GetAddress(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func GetAddressList(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getAddressList service.AddressService
	if err := c.ShouldBind(&getAddressList); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getAddressList.GetAddressList(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func DeleteAddress(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var deleteAddress service.AddressService
	if err := c.ShouldBind(&deleteAddress); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := deleteAddress.DeleteAddress(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func UpdateAddress(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var updateAddress service.AddressService
	if err := c.ShouldBind(&updateAddress); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := updateAddress.UpdateAddress(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
