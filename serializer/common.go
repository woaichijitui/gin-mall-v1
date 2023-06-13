package serializer

import (
	"gin-mall/pkg/e"
)

// 回复的通用模板
type Response struct {
	Status int         `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
	Msg    string      `json:"msg" form:"msg"`
	Error  string      `json:"error" form:"error"`
}

// ??
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// 轮播图回复模板
func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: e.Success,
		Data: DataList{
			item,
			total,
		},
		Msg: e.GetMsg(e.Success),
	}

}
