package service

import (
	"context"
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"` //商品种类
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"` //数量
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserByUId(uId)

	//以第一张图为封面
	tmp, _ := files[0].Open()
	path, err := uploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	path = conf.Host + conf.HttpPort + conf.ProductPath + path
	//将product 存入数据库
	product := model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(&product)
	if err != nil {
		code := e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	////将所有图片存入数据库
	//wg := new(sync.WaitGroup) // 创建
	//wg.Add(len(files))

	productImgDao := dao.NewProductImgDaoByDb(productDao.DB)

	for index, file := range files {
		num := strconv.Itoa(index)

		//多线程存入每个图片
		//go func(file *multipart.FileHeader) { //匿名函数

		//
		tmp, _ := file.Open()
		imgPath, err2 := uploadProductToLocalStatic(tmp, uId, service.Name+num) //将图片存入本地，并和封面图片path用num区分
		if err2 != nil {
			code = e.Error
		}
		defer tmp.Close()

		//将地址弄进数据库
		productImg := &model.ProductImg{
			ProductId: product.ID,
			ImgPath:   imgPath,
		}
		err2 = productImgDao.CreateProductImg(productImg)
		if err2 != nil {
			code = e.ErrorProductImgUpload
		}
		//wg.Done()
		//}(file)
	}
	//wg.Wait() //不加会报错
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}

func (service *ProductService) List(c context.Context) serializer.Response {
	var products []*model.Product
	var err error
	var code = e.Success

	//分页
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	//获取种类
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}

	productDao := dao.NewProductDao(c)
	//获取总数
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		//？？这个是重复操作吗？
		productDao = dao.NewProductDaoByDb(productDao.DB)
		products, _ = productDao.ListproductByCondition(condition, service.BasePage)
		wg.Done()

	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Show(c context.Context, id string) serializer.Response {
	var product *model.Product
	code := e.Success
	var err error

	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(c)
	product, err = productDao.GetProductById(uint(pId))
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
