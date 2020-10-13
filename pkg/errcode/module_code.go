package errcode

var (
	ErrorGetLinkListFail = NewError(20010001, "获取短链接列表失败")
	ErrorCreateLinkFail  = NewError(20010002, "创建短链接失败")
	ErrorUpdateLinkFail  = NewError(20010003, "更新短链接失败")
	ErrorDeleteLinkFail  = NewError(20010004, "删除短链接失败")
	ErrorCountLinkFail   = NewError(20010005, "统计短链接失败")
	ErrorUnShortLinkFail = NewError(20010006, "还原短链接失败")
	ErrorGetLinkFail     = NewError(20010007, "获取短链接失败")

	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
