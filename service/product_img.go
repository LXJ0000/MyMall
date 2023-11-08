package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/serializer"
	"context"
	"strconv"
)

type ProductImgService struct {
}

func (service *ProductImgService) ListProductImg(ctx context.Context, productId string) serializer.Response {
	code := e.Success
	pId, err := strconv.Atoi(productId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productImgDao := dao.NewProductImgDao(ctx)
	productImgList, err := productImgDao.ListProductImg(uint(pId))
	return serializer.BuildListResponse(serializer.BuildProductImgList(productImgList), uint(len(productImgList)))
}
