all: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/hello_linux_amd64 main.go

build-linux-arm64:
	mkdir -p build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/hello_linux_arm64 main.go