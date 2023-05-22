package dao

import (
	"fmt"
	"gin-mall/model"
)

func Migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Admin{},
			&model.Address{},
			&model.BasePage{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.Product{},
			&model.ProductImg{},
		)
	if err != nil {
		fmt.Printf("gorm migrate err :", err)

	}
	return
}
