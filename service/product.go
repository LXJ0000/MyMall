package service

import (
	"MyMall/pkg/e"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"MyMall/serializer"
	"context"
	"mime/multipart"
	"strconv"
	"sync"
	"time"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) CreateProduct(ctx context.Context, userId uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ := userDao.GetUserByUserId(userId)

	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        userId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	if err := productDao.CreateProduct(product); err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//	以第一张图片作为封面
	tmp, _ := files[0].Open()
	filePath, err := UploadToLocalStatic("product", product.ID, "top", files[0].Filename, tmp)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.ImgPath = filePath
	_ = productDao.UpdateProductImg(product)

	//	多图片创建 并发
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		//num := strconv.Itoa(idx)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		filePath, err = UploadToLocalStatic("product", product.ID, strconv.Itoa(int(time.Now().Unix())), file.Filename, tmp)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := &model.ProductImg{
			ProductId: product.ID,
			ImgPath:   filePath,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
