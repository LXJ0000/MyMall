package dao

import (
	"MyMall/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDbClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (n *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = n.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
