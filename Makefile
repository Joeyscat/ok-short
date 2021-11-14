.PHONY: all build clean default help init test format lint check-license
default: help

gomod:
	go mod tidy
	go mod vendor

build:
	go build -o ok-short cmd/ok-short/main.go

mockgen:
	mockgen --build_flags=--mod=mod -self_package=github.com/joeyscat/ok-short/internal/store -destination internal/store/mock_store.go -package store github.com/joeyscat/ok-short/internal/store Factory,LinkStore,LinkTraceStore
	mockgen --build_flags=--mod=mod -self_package=github.com/joeyscat/ok-short/internal/service/v1 -destination internal/service/v1/mock_service.go -package v1 github.com/joeyscat/ok-short/internal/service/v1 Service,LinkSrv,LinkTraceSrv

lint:
	golangci-lint run

test:
	go test -v ./...

clean:
	rm build/* -rf
