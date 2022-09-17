GOPATH:=$(shell go env GOPATH)

.PHONY: run build build-mac build-linux build-wasm build-windows test clean

run:
	@go run cmd/*.go

build: build-mac build-linux build-wasm build-windows

build-mac: clean
	@GOOS=darwin go build -trimpath -o ../dist/mac/zaid cmd/*.go

build-linux: clean
	@GOOS=linux go build -trimpath -o ../dist/linux/zaid cmd/*.go

build-wasm: clean
	@GOOS=js GOARCH=wasm go build -trimpath -o ../dist/wasm/zaid.wasm wasm/wasm.go

build-windows: clean
	@GOOS=windows go build -trimpath -o ../dist/windows/zaid.exe cmd/*.go

test:
	@go test -v -race -timeout 5s `go list ./... | grep -v "/wasm"` | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

clean:
	@rm -rf dist/mac
	@rm -rf dist/linux
	@rm -rf dist/wasm
	@rm -rf dist/windows
