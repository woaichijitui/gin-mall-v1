package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{
		NewDBClient(ctx), //绑定DB的上下文
	}
}
func NewProductImgDaoByDb(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{
		db,
	}
}

// 存入一条productImg
func (productImgDao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	err = productImgDao.DB.Model(&model.ProductImg{}).Create(productImg).Error
	return
}
