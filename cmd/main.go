package main

import (
	"gin-mall/conf"
	"gin-mall/routes"
)

func main() {
	conf.Init()

	r := routes.NewRouter()

	r.Run(":3000")
}
