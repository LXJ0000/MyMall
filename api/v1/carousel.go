package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.ListCarouselService
	if err := c.ShouldBind(&listCarousel); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err:", err)
		return
	}
	res := listCarousel.GetListCarousel(c.Request.Context())
	c.JSON(http.StatusOK, res)

}
