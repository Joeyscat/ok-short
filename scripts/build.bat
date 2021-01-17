SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ok-short cmd/ok-short/main.go
