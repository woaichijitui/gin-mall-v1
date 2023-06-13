package api

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(ctx *gin.Context) {
	var listCarousel service.CarouselService

	//绑定UserService
	if err := ctx.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
