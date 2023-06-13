package service

import (
	"context"
	"gin-mall/dao"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success

	//获取轮播map
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Info("err:", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))

}
