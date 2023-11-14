package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"strconv"
)

type CartService struct {
	ProductId uint `json:"product_id" form:"product_id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	Num       uint `json:"num" form:"num"`
	//MaxNum    uint `json:"max_num" form:"max_num"`
	//Check     bool `json:"check" form:"check"`
}

func (service *CartService) CreateCart(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	//判断商品是否存在
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "商品不存在",
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserByUserId(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "商家已跑路",
		}
	}
	cartDao := dao.NewCartDao(ctx)
	//判断商品是否在购物车中
	cart, err := cartDao.GetCartByProductAndUserId(service.ProductId, uId)
	if err != nil { //
		//不在购物车
		cart = &model.Cart{
			UserId:    uId,
			ProductId: service.ProductId,
			BossId:    service.BossId,
			Num:       1,
			MaxNum:    uint(product.Num), // 最大限购数量
			Check:     false,
		}
		_ = cartDao.CreateCart(cart)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildCart(cart, product, boss),
			Error:  "",
		}
	}
	if int(cart.Num) >= product.Num {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "加购失败，库存不足",
		}
	}
	cart.Num++
	err = cartDao.UpdateCartNumByCartIdAndUserId(cart.ID, uId, cart.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}
func (service *CartService) GetCart(ctx context.Context, uId uint, cartIdStr string) serializer.Response {
	code := e.Success
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	CartDao := dao.NewCartDao(ctx)
	cart, err := CartDao.GetCartByCartIdAndUserId(uint(cartId), uId)
	productDao := dao.NewProductDao(ctx)
	product, _ := productDao.GetProductById(cart.ProductId)
	bossDao := dao.NewUserDao(ctx)
	boss, _ := bossDao.GetUserByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}

}
func (service *CartService) GetCartList(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	cartList, err := CartDao.GetCartListByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCartList(ctx, cartList), uint(len(cartList)))

}
func (service *CartService) DeleteCart(ctx context.Context, uId uint, cartIdStr string) serializer.Response {
	code := e.Success
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	CartDao := dao.NewCartDao(ctx)
	_, err = CartDao.GetCartByCartIdAndUserId(uint(cartId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商品不存在",
		}
	}

	if err = CartDao.DeleteCartByCartIdAndUserId(uint(cartId), uId); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
func (service *CartService) UpdateCart(ctx context.Context, uId uint, cartIdStr string) serializer.Response {
	code := e.Success
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	CartDao := dao.NewCartDao(ctx)
	cart, _ := CartDao.GetCartByCartIdAndUserId(uint(cartId), uId)
	productDao := dao.NewProductDao(ctx)
	product, _ := productDao.GetProductById(cart.ProductId)
	bossDao := dao.NewUserDao(ctx)
	boss, _ := bossDao.GetUserByUserId(uId)
	if int(service.Num) > product.Num {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "库存不足",
		}
	}
	if err = CartDao.UpdateCartNumByCartIdAndUserId(uint(cartId), uId, service.Num); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cart.Num = service.Num
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}
