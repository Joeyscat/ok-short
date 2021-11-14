.PHONY: all build clean default help init test format lint check-license
default: help

build:
	go build -o ok-short cmd/ok-short/main.go

lint:
	golangci-lint run

clean:
	rm build/* -rf
