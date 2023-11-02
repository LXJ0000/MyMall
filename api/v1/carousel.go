package v1

import (
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.ListCarouselService
	if err := c.ShouldBind(&listCarousel); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	res := listCarousel.GetListCarousel(c.Request.Context())
	c.JSON(http.StatusOK, res)

}
