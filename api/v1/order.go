package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createOrder service.OrderService
	if err := c.ShouldBind(&createOrder); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := createOrder.CreateOrder(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func GetOrder(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getOrder service.OrderService
	if err := c.ShouldBind(&getOrder); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getOrder.GetOrder(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func GetOrderList(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var getOrderList service.OrderService
	if err := c.ShouldBind(&getOrderList); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := getOrderList.GetOrderList(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func DeleteOrder(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var deleteOrder service.OrderService
	if err := c.ShouldBind(&deleteOrder); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := deleteOrder.DeleteOrder(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
