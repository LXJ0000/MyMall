package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Category      uint   `json:"category"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(product *model.Product) *Product {
	return &Product{
		Id:            product.ID,
		Name:          product.Name,
		Category:      product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.View(),
		CreatedAt:     product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossId:        product.BossId,
		BossName:      product.BossName,
		BossAvatar:    config.Host + config.HttpPort + config.AvatarPath + product.BossAvatar,
	}
}

func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, *product)
	}
	return products
}
