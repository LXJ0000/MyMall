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

		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		//商品
		//获取商品列表
		v1.GET("product", api.ListProduct)
		//搜索商品
		v1.POST("products", api.SearchProduct)
		//展示商品信息
		v1.GET("product/:id", api.ShowProduct)
		//展示商品图片
		v1.GET("img/:id", api.ListProductImg)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.PUT("avatar", api.UserUploadAvatar)
			authed.POST("user/sending-email", api.UserSendingEmail)
			authed.POST("user/valid-email", api.UserValidEmail)

			//显示金额
			authed.POST("money", api.ShowUserMoney)

			//商品操作
			//创建商品
			authed.POST("product", api.CreateProduct)

		}

	}

	return r
}
