package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB //为什么可以这样用
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{
		NewDBClient(ctx), //绑定DB的上下文
	}
}
func NewCarouselDaoByDb(db *gorm.DB) *CarouselDao {
	return &CarouselDao{
		db,
	}
}

// 根据id 获取Carousel
func (carouselDao *CarouselDao) GetCarouselDaoUId(id uint) (carousel *model.Carousel, err error) {

	err = carouselDao.DB.Model(&model.Carousel{}).Where("id = ?", id).First(&carousel).Error
	return
}

// ListCarousel 获取轮播list
func (carouselDao *CarouselDao) ListCarousel() (carousels []model.Carousel, err error) {
	err = carouselDao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
