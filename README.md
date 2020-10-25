## 短链接服务后台

### 项目结构
```bash
├── configs # 配置
│   └── config.yaml
├── docs # 文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── global
│   ├── db.go
│   ├── redis.go
│   └── setting.go
├── go.mod
├── go.sum
├── internel
│   ├── dao
│   │   ├── auth.go
│   │   ├── dao.go
│   │   ├── link.go
│   │   └── link_trace.go
│   ├── middleware
│   │   └── translations.go
│   ├── model
│   │   ├── auth.go
│   │   ├── link.go
│   │   ├── link_trace.go
│   │   ├── model.go
│   │   └── README.md
│   ├── routers
│   │   ├── api
│   │   │   ├── auth.go
│   │   │   └── v1
│   │   │       ├── link.go
│   │   │       └── link_trace.go
│   │   └── router.go
│   └── service
│       ├── auth.go
│       ├── link.go
│       ├── link_trace.go
│       └── service.go
├── LICENSE
├── main.go
├── pkg
│   ├── app
│   │   ├── app.go
│   │   ├── form.go
│   │   ├── jwt.go
│   │   ├── pagination.go
│   │   └── utils.go
│   ├── convert
│   │   └── convert.go
│   ├── errcode
│   │   ├── common_code.go
│   │   ├── errcode.go
│   │   └── module_code.go
│   ├── logger
│   │   └── logger.go
│   ├── setting
│   │   ├── section.go
│   │   └── setting.go
│   └── util
│       └── md5.go
├── README.md
├── scripts # 脚本
│   └── db.sql
└── vendor # 第三方库
```