package errcode

var (
	Success                  = NewError(0, "成功")
	NotFound                 = NewError(100401, "找不到")
	InvalidParams            = NewError(100402, "参数错误，请检查参数")
	UnauthorizedAuthNotExist = NewError(100403, "找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError   = NewError(100404, "Token错误")
	UnauthorizedTokenTimeout = NewError(100405, "Token超时")
	UnauthorizedLoginFail    = NewError(100406, "登录失败，用户不存在或密码不正确")
	TooManyRequests          = NewError(100407, "请求过多")
	ServerError              = NewError(100500, "服务内部错误")
	TokenGenerateError       = NewError(100501, "Token生成失败")
)
