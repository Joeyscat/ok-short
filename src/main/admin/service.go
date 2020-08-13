package admin

type Service struct {
	// 用户管理
	CreatorService
	// 短链管理
	LinkService
	// 短链访问管理
	VisitorService
}
