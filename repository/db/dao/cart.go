package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDbClient(ctx)}
}
func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}
func (dao *CartDao) CreateCart(cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(cart).Error
}
func (dao *CartDao) GetCartByCartIdAndUserId(aId uint, uId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", aId, uId).First(&cart).Error
	return
}
func (dao *CartDao) GetCartListByUserId(uId uint) (list []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&list).Error
	//todo 分页
	return
}
func (dao *CartDao) GetCartByProductAndUserId(pId uint, uId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("product_id=? AND user_id=?", pId, uId).First(&cart).Error
	return
}
func (dao *CartDao) DeleteCartByCartIdAndUserId(Id uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", Id, uId).Delete(&model.Cart{}).Error
	return
}

// UpdateCartNumByCartIdAndUserId
//
//	func (dao *CartDao) UpdateCartByCartIdAndUserId(Id uint, uId uint, cart *model.Cart) (err error) {
//		err = dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", Id, uId).Updates(&cart).Error
//		return
//	}
func (dao *CartDao) UpdateCartNumByCartIdAndUserId(Id uint, uId uint, num uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=? AND user_id=?", Id, uId).Update("num", num).Error
	return
}
