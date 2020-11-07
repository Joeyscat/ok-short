### ok-short
基于golang构建的短链接服务

### 技术栈
* gin
* gorm
* redis
* mysql
* swagger

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## 功能
* 长链接->短链接
* 短链接跳转
* 短链接查询

### 运行方式
1. 编辑配置文件 configs/config.yaml
主要是 MySQL 和 Redis 配置，以及 App.LinkPrefix，这是短链服务所在主机的域名，短链接就是根据这个域名和一个短码生成的
2. 编译项目&运行
```shell script
go build -o app
./app
```
3. 访问swagger http://ip:port/swagger/index.html

### 许可
Released under the [MIT License](https://github.com/Joeyscat/ok-short/blob/master/LICENSE)