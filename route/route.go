package route

import (
	api "MyMall/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})

		v1.POST("user/register", api.UserRegister)
	}

	return r
}
