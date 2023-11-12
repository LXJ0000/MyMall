package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDbClient(ctx)}
}
func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}
func (dao *CategoryDao) GetCategory() (categoryList []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categoryList).Error
	return
}
