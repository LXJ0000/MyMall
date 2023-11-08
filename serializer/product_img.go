package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) *ProductImg {
	return &ProductImg{
		ProductId: item.ProductId,
		ImgPath:   config.Host + config.HttpPort + config.ProductPath + item.ImgPath,
	}
}

func BuildProductImgList(items []*model.ProductImg) (productImgList []*ProductImg) {
	for _, item := range items {
		productImgList = append(productImgList, BuildProductImg(item))
	}
	return
}
