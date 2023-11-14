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
		//展示商品分类
		v1.GET("categories", api.ListCategories)

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

			//	收藏夹操作
			authed.POST("favorite", api.CreateFavorite)
			authed.GET("favorite", api.ShowFavorite)
			authed.DELETE("favorite/:id", api.DeleteFavorite)

			//	地址操作
			authed.POST("address", api.CreateAddress)
			authed.GET("address", api.GetAddressList)
			authed.GET("address/:id", api.GetAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

			//	购物车操作
			authed.POST("cart", api.CreateCart)
			authed.DELETE("cart/:id", api.DeleteCart)
			authed.GET("cart/:id", api.GetCart)
			authed.PUT("cart/:id", api.UpdateCart)
			authed.GET("cart", api.GetCartList)

			//	订单操作
			authed.POST("order", api.CreateOrder)
			authed.GET("order/:id", api.GetOrder)
			authed.DELETE("order/:id", api.DeleteOrder)
			authed.GET("order", api.GetOrderList)
		}

	}

	return r
}
