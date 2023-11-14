package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pay(c *gin.Context) {
	pay := service.PayService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&pay); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	res := pay.Pay(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}
