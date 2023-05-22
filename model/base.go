package model

import "gorm.io/gorm"

type BasePage struct {
	gorm.Model
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
