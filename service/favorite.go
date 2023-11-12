package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) CreateFavorite(ctx context.Context, userId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	if exist := favoriteDao.FavoriteExistOrNot(service.ProductId, userId); exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByUserId(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserByUserId(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserId:    userId,
		Product:   *product,
		ProductId: service.ProductId,
		Boss:      *boss,
		BossId:    service.BossId,
	}
	if err := favoriteDao.CreateFavorite(favorite); err != nil {
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
		Data:   serializer.BuildFavorite(favorite, product, boss),
	}
}
func (service *FavoriteService) ShowFavorite(ctx context.Context, userId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	favoriteList, err := favoriteDao.ListFavorite(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavoriteList(ctx, favoriteList), uint(len(favoriteList)))
}
func (service *FavoriteService) DeleteFavorite(ctx context.Context, userId uint, favoriteIdStr string) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	fId, err := strconv.Atoi(favoriteIdStr)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := favoriteDao.DeleteFavorite(userId, uint(fId)); err != nil {
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
