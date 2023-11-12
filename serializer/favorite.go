package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"context"
)

type Favorite struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	ProductName   string `json:"product_name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DisCountPrice string `json:"dis_count_price"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) *Favorite {
	return &Favorite{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		CreatedAt:     favorite.CreatedAt.Unix(),
		ProductName:   product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
		Price:         product.Price,
		DisCountPrice: product.DiscountPrice,
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossId:        boss.ID,
	}
}

func BuildFavoriteList(ctx context.Context, items []*model.Favorite) (FavoriteList []*Favorite) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, v := range items {
		product, _ := productDao.GetProductById(v.ProductId)
		boss, _ := bossDao.GetUserByUserId(v.UserId)
		FavoriteList = append(FavoriteList, BuildFavorite(v, product, boss))
	}
	return
}
