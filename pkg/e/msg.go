package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExitUser:              "存在该用户",
	ErrorFailEncry:             "密码加密错误",
	ErrorUserNotFount:          "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token签发失败",
	ErrorAuthCheckTokenTimeOut: "token过期",
	ErrorUploadFail:            "上传失败",
}

// 获取msg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
