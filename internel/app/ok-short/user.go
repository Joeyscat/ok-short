package ok_short

type UserService struct {
}

// 账号注册
func (*UserService) Register(account, password string) {

}

// 通过第三方平台注册(微信、QQ、钉钉)
func (*UserService) RegisterBy(account, accessToken string) {

}

// 获取用户登录专用的token，该token具有一次性、实效性
func (*UserService) LoginToken(account string) (string, error) {

	return "", nil
}

// 用户账户密码登录，password = hash(hash(真实密码)+token)
func (u *UserService) Login(account, password, token string) (string, error) {
	// 检测token有效性，并删除

	// 根据account查询password

	// 确认密码哈希值是否正确

	// 生成token并存入redis

	// 返回用户token
	return "", nil

	// TODO 这个token能干什么，泄露了会有什么风险
}

// 用户登出，删除redis中的token
func (*UserService) Logout(token string) {

}

func (*UserService) UserInfo(token string) string {

	return ""
}
