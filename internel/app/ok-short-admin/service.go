package ok_short_admin

type Service struct {
	// 用户管理
	AuthorService
	// 管理员
	UserService
	// 短链管理
	LinkService
}
