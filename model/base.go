package model

import "gorm.io/gorm"

type BasePage struct {
	gorm.Model
	PageNum  int `form:"page_num"`  //第几页
	PageSize int `form:"page_size"` //每页多少个
}
