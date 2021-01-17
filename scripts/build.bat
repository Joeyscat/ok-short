SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ok-short.exe cmd/ok-short/main.go
