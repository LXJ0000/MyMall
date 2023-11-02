package route

import (
	api "MyMall/api/v1"
	"MyMall/middleware"
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

		//	用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			//	用户操作
			authed.PUT("user", api.UserUpdate)
			authed.PUT("avatar", api.UserUploadAvatar)
			authed.POST("user/sending-email", api.UserSendingEmail)
			authed.POST("user/valid-email", api.UserValidEmail)

			//	显示金额
			authed.POST("money", api.ShowUserMoney)

		}

	}

	return r
}
