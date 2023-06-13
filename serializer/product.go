package serializer

import (
	"gin-mall/model"
)

type Product struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	View          uint64 `json:"view"` //点击人数
	CreateAt      int64  `json:"create_at"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		OnSale:        item.OnSale,
		Num:           item.Num,
		View:          item.View(),
		CreateAt:      item.CreatedAt.Unix(),
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    item.BossAvatar,
	}
}

// 展示每页商品模板
func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return
}
