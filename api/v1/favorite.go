package v1

import (
	util "MyMall/pkg/utils"
	"MyMall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorite(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var createFavoriteService service.FavoriteService
	if err := c.ShouldBind(&createFavoriteService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := createFavoriteService.CreateFavorite(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}

func ShowFavorite(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var showFavoriteService service.FavoriteService
	if err := c.ShouldBind(&showFavoriteService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := showFavoriteService.ShowFavorite(c.Request.Context(), claims.ID)
	c.JSON(http.StatusOK, res)
}
func DeleteFavorite(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	var deleteFavoriteService service.FavoriteService
	if err := c.ShouldBind(&deleteFavoriteService); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Errorln("err", err)
		return
	}
	res := deleteFavoriteService.DeleteFavorite(c.Request.Context(), claims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
