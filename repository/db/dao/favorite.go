package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDbClient(ctx)}
}
func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}
func (dao *FavoriteDao) ListFavorite(uId uint) (listFavorite []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", uId).Find(&listFavorite).Error
	return
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoriteDao) FavoriteExistOrNot(pId uint, uId uint) bool {
	err := dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND product_id = ?", uId, pId).First(&model.Favorite{}).Error
	return err == nil
}
func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).Where("id=? AND user_id=?", fId, uId).Delete(&model.Favorite{}).Error
}
