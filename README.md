## 项目结构
```bash
.
├── api # 接口/协议
│   ├── link.go
│   ├── link.pb.go
│   ├── link.proto
│   ├── link_test.go
│   └── plugin.go
├── cmd # 程序入口
│   └── ok-short
│       ├── main.go
│       └── ok-short
├── go.mod
├── go.sum
├── internel # 内部代码
│   ├── app
│   │   ├── app.go
│   │   ├── context.go
│   │   ├── ok-short # 服务实现
│   │   │   ├── link.go
│   │   │   ├── link_test.go
│   │   │   └── service.go
│   │   ├── ok-short-admin # 后台管理
│   │   │   ├── author.go
│   │   │   ├── link.go
│   │   │   ├── service.go
│   │   │   ├── user.go
│   │   │   └── user_test.go
│   │   └── route_handler.go
│   └── pkg # 通用代码
│       ├── common
│       │   ├── error.go
│       │   ├── status.go
│       │   ├── utils.go
│       │   └── view_model.go
│       ├── middleware.go
│       ├── model
│       │   ├── link.go
│       │   └── user.go
│       └── store.go
├── README.md
├── scripts # 脚本
│   └── db.sql
└── vendor # 第三方库
```