package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDbClient(ctx)}
}
func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
func (dao *ProductDao) UpdateProductImg(id uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id=?", id).Updates(&product).Error
}
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, basePage *model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Offset((basePage.PageNum - 1) * basePage.PageSize).Limit(basePage.PageSize).Find(&products).Error
	return
}
func (dao *ProductDao) CountProductByInfo(info string) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ? OR name LIKE ?", "%"+info+"%", "%"+info+"%", "%"+info+"%").Count(&total).Error
	return
}
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? OR info LIKE ? OR name LIKE ?", "%"+info+"%", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&products).
		Error
	return
}
func (dao *ProductDao) GetProductById(pId uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", pId).First(&product).Error
	return
}
