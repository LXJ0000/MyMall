package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createProductService service.ProductService
	if err := c.ShouldBind(&createProductService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := createProductService.CreateProduct(c.Request.Context(), claims.ID, files)
	c.JSON(http.StatusOK, res)
}
