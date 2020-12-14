# docker docker build -t ok-short . && docker run ok-short
# docker build --build-arg app=async-db -t async-db . && docker run async-db

FROM golang:alpine AS builder

ARG app=ok-short

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件
RUN go build -o app cmd/$app/main.go

###################
# 接下来创建一个小镜像
###################
FROM scratch

COPY ./configs /configs

# 从builder镜像中把/build/app 拷贝当前目录
COPY --from=builder /build/app /

EXPOSE 8700

# 需要运行的命令
CMD ["/app"]