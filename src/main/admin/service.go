package admin

type Service struct {
	// 用户管理
	AuthorService
	// 管理员
	UserService
	// 短链管理
	LinkService
	// 短链访问管理
	VisitorService
}
