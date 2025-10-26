package code

type Errno struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	OK                  = Errno{1000, "成功"}
	UserExists          = Errno{1001, "用户已存在"}
	InternalServerError = Errno{1002, "内部服务器错误"}
	UserNotFound        = Errno{1003, "用户不存在"}
	InvalidPassword     = Errno{1004, "密码错误"}
	InvalidToken        = Errno{1005, "无效的Token"}
	TokenExpired        = Errno{1006, "Token已过期"}
	TokenMalformed      = Errno{1007, "Token格式错误"}
	InvalidParams       = Errno{1008, "无效的参数"}
	NotFound            = Errno{1009, "资源不存在"}
)
