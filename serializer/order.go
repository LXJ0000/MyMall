package serializer

import (
	"MyMall/config"
	"MyMall/repository/db/dao"
	"MyMall/repository/db/model"
	"context"
)

type Order struct {
	OrderId      uint    `json:"order_id"`
	OrderNum     string  `json:"order_num"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	UserId       uint    `json:"user_id"`
	BossId       uint    `json:"boss_id"`
	ProductId    uint    `json:"product_id"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"money"`
}

func BuildOrder(order *model.Order, addr *model.Address, product *model.Product) *Order {
	return &Order{
		OrderId:   order.ID,
		OrderNum:  order.OrderNum,
		CreatedAt: order.CreatedAt.Unix(),
		UpdatedAt: order.UpdatedAt.Unix(),
		UserId:    order.UserId,
		BossId:    order.BossId,
		ProductId: order.ProductId,
		Type:      order.Type,
		Money:     order.Money,

		AddressName:  addr.Name,
		AddressPhone: addr.Phone,
		Address:      addr.Address,
		ProductName:  product.Name,
		ImgPath:      config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
	}
}

func BuildOrderList(ctx context.Context, Orders []*model.Order) (list []*Order) {
	productDao := dao.NewProductDao(ctx)
	addrDao := dao.NewAddressDao(ctx)
	for _, item := range Orders {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		addr, err := addrDao.GetAddressByAddressIdAndUserId(item.AddressId, item.UserId)
		if err != nil {
			continue
		}
		list = append(list, BuildOrder(item, addr, product))
	}
	return
}
