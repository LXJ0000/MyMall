package service

import (
	"MyMall/pkg/e"
	util "MyMall/pkg/utils"
	"MyMall/repository/db/dao"
	"MyMall/serializer"
	"context"
)

type ListCarouselService struct {
}

func (l *ListCarouselService) GetListCarousel(ctx context.Context) serializer.Response {
	code := e.Success
	carouselDao := dao.NewCarouselDao(ctx)
	carousel, err := carouselDao.GetCarousel()
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarouselList(carousel), uint(len(carousel)))
}
