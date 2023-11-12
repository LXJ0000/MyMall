package service

import (
	"MyMall/pkg/e"
	util "MyMall/pkg/utils"
	"MyMall/repository/db/dao"
	"MyMall/serializer"
	"context"
)

type CategoryService struct {
}

func (service *CategoryService) ListCategories(ctx context.Context) serializer.Response {
	code := e.Success
	categoryDao := dao.NewCategoryDao(ctx)
	category, err := categoryDao.GetCategory()
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategoryList(category), uint(len(category)))
}
