package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDbClient(ctx)}
}
func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}
func (dao *OrderDao) CreateOrder(Order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(Order).Error
}
func (dao *OrderDao) GetOrderByOrderIdAndUserId(aId uint, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", aId, uId).First(&order).Error
	return
}
func (dao *OrderDao) GetOrderListByUserId(uId uint) (list []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id=?", uId).Find(&list).Error
	return
}
func (dao *OrderDao) GetOrderByProductAndUserId(pId uint, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("product_id=? AND user_id=?", pId, uId).First(&order).Error
	return
}
func (dao *OrderDao) DeleteOrderByOrderIdAndUserId(Id uint, uId uint) (err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", Id, uId).Delete(&model.Order{}).Error
	return
}
func (dao *OrderDao) GetOrderByCondition(condition map[string]interface{}, page *model.BasePage) (list []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageNum).
		Limit(page.PageSize).
		Find(&list).
		Error

	return
}
func (dao *OrderDao) GetOrderCountByCondition(condition map[string]interface{}) (cnt int64) {
	dao.DB.Model(&model.Order{}).Where(condition).Find(&model.Order{}).Count(&cnt)
	return
}
func (dao *OrderDao) UpdateOrderTypeById(id uint, TYPE uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", id).Update("type", TYPE).Error
}
