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

func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
func (dao *ProductDao) UpdateProductImg(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Updates(product).Error
}
