
```bash
└── src
    └── main
        ├── admin # 后台管理相关
        │   ├── creator_service.go
        │   ├── link_service.go
        │   ├── README.md
        │   ├── service.go
        │   └── visitor_service.go
        ├── api # 短链服务相关
        │   ├── link_service.go
        │   ├── link_service_test.go
        │   └── service.go
        ├── app.go
        ├── common # 通用代码
        │   ├── error.go
        │   ├── link.go
        │   └── utils.go
        ├── env.go
        ├── go.mod
        ├── go.sum
        ├── main.go # 程序入口
        ├── middleware.go
        ├── route_handler.go
        ├── shorten_view_model.go
        └── store # 存储相关
            ├── mysql.go
            ├── mysql_test.go
            └── redis.go

```