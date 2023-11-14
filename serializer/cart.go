package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"context"
)

type Cart struct {
	CartId         uint   `json:"Cart_id"`
	UserId         uint   `json:"user_id"`
	ProductId      uint   `json:"product_id"`
	BossId         uint   `json:"boss_id"`
	BossName       string `json:"boss_name"`
	CreatedAt      int    `json:"created_at"`
	Num            uint   `json:"num"`
	MaxNum         uint   `json:"max_num"`
	Check          bool   `json:"check"`
	ProductName    string `json:"product_name"`
	ProductImgPath string `json:"product_img_path"`
	DiscountPrice  string `json:"discount_price"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) *Cart {
	return &Cart{
		CartId:         cart.ID,
		UserId:         cart.UserId,
		ProductId:      cart.ProductId,
		BossId:         cart.BossId,
		BossName:       boss.UserName,
		CreatedAt:      int(cart.CreatedAt.Unix()),
		Num:            cart.Num,
		MaxNum:         uint(product.Num),
		Check:          cart.Check,
		ProductName:    product.Name,
		ProductImgPath: config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
		DiscountPrice:  product.DiscountPrice,
	}
}

func BuildCartList(ctx context.Context, carts []*model.Cart) (list []*Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range carts {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserByUserId(item.UserId)
		if err != nil {
			continue
		}
		list = append(list, BuildCart(item, product, boss))
	}
	return
}
