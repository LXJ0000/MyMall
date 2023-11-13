package serializer

import "MyMall/repository/db/model"

type Address struct {
	AddressId uint   `json:"address_id"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func BuildAddress(addr *model.Address) *Address {
	return &Address{
		AddressId: addr.ID,
		UserId:    addr.UserId,
		Name:      addr.Name,
		Phone:     addr.Phone,
		Address:   addr.Address,
	}
}

func BuildAddressList(addr []*model.Address) (list []*Address) {
	for _, item := range addr {
		list = append(list, BuildAddress(item))
	}
	return
}
