package serializer

import "gin-mall/model"

// 将dao获取到的carousel map 序列化输出
type Carousel struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

// 创建一条序列化输出
func BuildCarousel(item *model.Carousel) Carousel {

	return Carousel{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

// 创建将所有序列化输出
func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(&item)
		carousels = append(carousels, carousel)
	}
	return
}
