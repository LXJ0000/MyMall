package service

import (
	"MyMall/pkg/e"
	util "MyMall/pkg/utils"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"fmt"
	"strconv"
)

type PayService struct {
	OrderId uint `json:"order_id" form:"order_id"`
	//Money     float64 `json:"money" form:"money"`
	//OrderNum string `json:"order_num" form:"order_num"`
	//ProductId uint    `json:"product_id" form:"product_id"`
	//PayTime   uint    `json:"pay_time" form:"pay_time"` // 支付时间
	//Sign      string  `json:"sign" form:"sign"`         // 信号
	//BossId    uint    `json:"boss_id" form:"boss_id"`
	//BossName  string  `json:"boss_name" form:"boss_name"`
	Key string `json:"key" form:"key"` // 支付金额？？？
	//Num       int     `json:"num" form:"num"`
}

func (service *PayService) Pay(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderByOrderIdAndUserId(service.OrderId, uId)
	if order.Type == 2 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "请勿重复支付",
		}

	}

	util.Encrypt.SetKey(service.Key)
	tx := orderDao.Begin()
	money := order.Money * float64(order.Num)
	userDao := dao.NewUserDao(ctx)
	user, _ := userDao.GetUserByUserId(uId)

	//	对钱解密，减去money，再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "余额不足",
		}
	}
	//y用户扣钱
	finalMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finalMoney)
	userDao = dao.NewUserDaoByDB(userDao.DB)
	if err = userDao.UpdateUserByUserId(uId, user); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "更新用户信息失败",
		}
	}
	boss, err := userDao.GetUserByUserId(order.BossId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商家以跑路",
		}
	}
	//商家加钱
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finalMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finalMoney)

	if err = userDao.UpdateUserByUserId(order.BossId, boss); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "更新商家信息失败",
		}
	}
	//	商品数 - 1
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "商品不存在",
		}
	}
	product.Num -= order.Num
	if productDao.UpdateProductById(product.ID, product) != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "更新商品信息失败",
		}
	}
	if orderDao.UpdateOrderTypeById(order.ID, 2) != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "订单类型更新失败",
		}
	}
	//	自己的商品 + 1 同一件商品？
	UserProduct, err := productDao.GetProductByProductIdAndBossId(product.ID, uId)
	if err != nil {
		//	没有该商品
		UserProduct = &model.Product{
			Name:          product.Name,
			CategoryId:    product.CategoryId,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			OnSale:        false,
			Num:           1,
			BossId:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}

		if productDao.CreateProduct(UserProduct) != nil {
			tx.Rollback()
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  "购买商品失败",
			}
		}
	} else {
		UserProduct.Num += order.Num
		if productDao.UpdateProductById(UserProduct.ID, UserProduct) != nil {
			tx.Rollback()
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  "修改拥有商品数失败",
			}
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
