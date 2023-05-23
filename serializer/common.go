package serializer

type Response struct {
	Status int         `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
	Msg    string      `json:"msg" form:"msg"`
	Error  string      `json:"error" form:"error"`
}
