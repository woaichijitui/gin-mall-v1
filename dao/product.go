package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{
		NewDBClient(ctx), //绑定DB的上下文
	}
}
func NewProductDaoByDb(db *gorm.DB) *ProductDao {
	return &ProductDao{
		db,
	}
}

// 创建一个product
func (productDao *ProductDao) CreateProduct(product *model.Product) (err error) {
	err = productDao.DB.Model(&model.Product{}).Create(product).Error
	return
}

// 根据id 获取notice
func (productDao *ProductDao) GetProductByUId(id uint) (product *model.Product, err error) {

	err = productDao.DB.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	return
}

func (productDao *ProductDao) CountProductByCondition(condition map[string]interface{}) (count int64, err error) {
	err = productDao.DB.Model(&model.Product{}).Where(condition).Count(&count).Error
	return
}
func (productDao *ProductDao) ListproductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = productDao.DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

func (productDao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = productDao.DB.Model(&model.Product{}).Where("id=?", id).
		First(&product).Error //此处product 要用指针
	return
}
