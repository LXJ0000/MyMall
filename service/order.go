package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint   `json:"product_id" form:"product_id"`
	Num       int    `json:"num" form:"num"`
	AddressId uint   `json:"address_id" form:"address_id"`
	Type      uint   `json:"type" form:"type"`
	OrderNum  string `json:"order_num" form:"order_num"`
	model.BasePage
}

func (service *OrderService) CreateOrder(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商品不存在",
		}
	}
	//校验地址是否存在
	addrDao := dao.NewAddressDao(ctx)
	addr, err := addrDao.GetAddressByAddressIdAndUserId(service.AddressId, uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "地址信息有误",
		}
	}

	orderDao := dao.NewOrderDao(ctx)
	productDiscountPrice, _ := strconv.ParseFloat(product.DiscountPrice, 64)
	order := &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    product.BossId,
		AddressId: service.AddressId,
		Type:      1, // 默认未支付 = 1
		Money:     float64(service.Num) * productDiscountPrice,
		Num:       service.Num,
	}

	//创建订单号
	//生成9位随机数字
	//seed time.Now().UnixNano()
	//rand.NewSource(seed)
	//rand.New(源)
	//Int31n(n) 0~n int32
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	order.OrderNum = number + strconv.Itoa(int(service.ProductId)) + strconv.Itoa(int(uId))

	err = orderDao.CreateOrder(order)
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
		Data:   serializer.BuildOrder(order, addr, product),
	}
}
func (service *OrderService) GetOrder(ctx context.Context, uId uint, OrderIdStr string) serializer.Response {
	code := e.Success
	OrderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	OrderDao := dao.NewOrderDao(ctx)
	Order, err := OrderDao.GetOrderByOrderIdAndUserId(uint(OrderId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商品不存在",
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, _ := productDao.GetProductById(Order.ProductId)
	addrDao := dao.NewAddressDao(ctx)
	addr, _ := addrDao.GetAddressByAddressIdAndUserId(service.AddressId, uId)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(Order, addr, product),
	}

}
func (service *OrderService) GetOrderList(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	OrderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	if service.Type != 0 { // 查询全部订单
		condition["type"] = service.Type
	}
	condition["user_id"] = uId

	OrderList, err := OrderDao.GetOrderByCondition(condition, &service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrderList(ctx, OrderList), uint(OrderDao.GetOrderCountByCondition(condition)))
}

func (service *OrderService) DeleteOrder(ctx context.Context, uId uint, OrderIdStr string) serializer.Response {
	code := e.Success
	OrderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	OrderDao := dao.NewOrderDao(ctx)
	if err = OrderDao.DeleteOrderByOrderIdAndUserId(uint(OrderId), uId); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "订单不存在",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
