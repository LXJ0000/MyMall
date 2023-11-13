package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDbClient(ctx)}
}
func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}
func (dao *AddressDao) CreateAddress(addr *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(addr).Error
}
func (dao *AddressDao) GetAddressByAddressIdAndUserId(aId uint, uId uint) (addr *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).First(&addr).Error
	return
}
func (dao *AddressDao) GetAddressListByUserId(uId uint) (list []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&list).Error
	return
}
func (dao *AddressDao) DeleteAddressByAddressIdAndUserId(aId uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).Delete(&model.Address{}).Error
	return
}
func (dao *AddressDao) UpdateAddressByAddressIdAndUserId(aId uint, uId uint, addr *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).Updates(&addr).Error
	return
}
