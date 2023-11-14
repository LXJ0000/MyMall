package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createCart service.CartService
	if err := c.ShouldBind(&createCart); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := createCart.CreateCart(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func GetCart(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getCart service.CartService
	if err := c.ShouldBind(&getCart); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getCart.GetCart(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func GetCartList(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getCartList service.CartService
	if err := c.ShouldBind(&getCartList); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getCartList.GetCartList(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func DeleteCart(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var deleteCart service.CartService
	if err := c.ShouldBind(&deleteCart); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := deleteCart.DeleteCart(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func UpdateCart(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var updateCart service.CartService
	if err := c.ShouldBind(&updateCart); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := updateCart.UpdateCart(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
