package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type NoticDao struct {
	*gorm.DB
}

func NewNoticDao(ctx context.Context) *NoticDao {
	return &NoticDao{
		NewDBClient(ctx), //绑定DB的上下文
	}
}
func NewNoticDaoByDb(db *gorm.DB) *NoticDao {
	return &NoticDao{
		db,
	}
}

// 根据id 获取notice
func (noticeDao *NoticDao) GetNoticeByUId(id uint) (notice *model.Notice, err error) {

	err = noticeDao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return
}
