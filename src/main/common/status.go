package common

const (
	SystemBusy = -2
	SystemErr  = -1
	Success    = 0
	// Token 1000~1999
	TokenInvalid = 1000
	TokenExpired = 1001
	// User 2000~2999
	UserNotExists     = 2000
	UserAlreadyExists = 2001
	UserRegisterFail  = 2002
	UserLoginFail     = 2003
	UserAlreadyLogin  = 2004
	// Param 3000~3999
	ParamIllegal         = 3000
	ParamUserNameEmpty   = 3001
	ParamPasswordEmpty   = 3002
	ParamTokenEmpty      = 3003
	ParamURLEmpty        = 3004
	ParamURLInvalid      = 3005
	ParamPostBodyInvalid = 3006
	// Link 4000~4999
	LinkInvalid    = 4000
	LinkNotExists  = 4001
	LinkCreateFail = 4002

	ApiUnAuthorized = 48001
	EmptyPostData   = 44002
)

var statusText = map[int]string{
	SystemBusy:           "系统繁忙，请稍候再试",
	SystemErr:            "系统错误，请联系管理员",
	Success:              "请求成功",
	TokenInvalid:         "Token无效",
	TokenExpired:         "Token已过期",
	UserNotExists:        "用户不存在",
	UserAlreadyExists:    "用户已存在",
	UserRegisterFail:     "用户注册失败",
	UserLoginFail:        "用户登录失败",
	UserAlreadyLogin:     "用户已在别处登录",
	ParamIllegal:         "参数不合法",
	ParamUserNameEmpty:   "参数[用户名]为空",
	ParamPasswordEmpty:   "参数[密码]为空",
	ParamTokenEmpty:      "参数[Token]为空",
	ParamURLEmpty:        "参数[URL]为空",
	ParamURLInvalid:      "参数[URL]不合法",
	ParamPostBodyInvalid: "Post请求参数不合法",
	LinkInvalid:          "短链接无效",
	LinkNotExists:        "短链接不存在",
	LinkCreateFail:       "新增短链数据失败",
}

func BSText(code int) string {
	return statusText[code]
}
