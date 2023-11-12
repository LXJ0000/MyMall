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

func ListProduct(c *gin.Context) {
	var listProductService service.ProductService
	if err := c.ShouldBind(&listProductService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := listProductService.ListProduct(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func SearchProduct(c *gin.Context) {
	var searchProduct service.ProductService
	if err := c.ShouldBind(&searchProduct); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := searchProduct.SearchProduct(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func ShowProduct(c *gin.Context) {
	var showProduct service.ProductService
	if err := c.ShouldBind(&showProduct); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := showProduct.ShowProduct(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListProductImg(c *gin.Context) {
	var listProductImg service.ProductImgService
	if err := c.ShouldBind(&listProductImg); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := listProductImg.ListProductImg(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListCategories(c *gin.Context) {
	var listCategories service.CategoryService
	if err := c.ShouldBind(&listCategories); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := listCategories.ListCategories(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
