package model

import (
	"fmt"
	"gin-mall/cache"
	"gin-mall/pkg/util"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

// ?这是什么
func (product *Product) View() uint64 {
	_, err2 := fmt.Println("product.ID :", cache.RedisCtx)
	if err2 != nil {

		util.LogrusObj.Trace(err2)
	}
	countStr, _ := cache.RedisClient.Get(cache.RedisCtx, cache.ProductViemKey(product.ID)).Result()

	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// ?这是什么
func (product *Product) AddView() {
	cache.RedisClient.Incr(cache.RedisCtx, cache.ProductViemKey(product.ID)) //incr:增长
	cache.RedisClient.ZIncrBy(cache.RedisCtx, cache.RangKey, 1, strconv.Itoa(int(product.ID)))
}
