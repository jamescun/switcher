.PHONY: build release

build:
	go build -o switcher *.go

release:
	GOOS=linux GOARCH=amd64 go build -o switcher-linux-x64
	GOOS=darwin GOARCH=amd64 go build -o switcher-darwin-x64
	GOOS=windows GOARCH=amd64 go build -o switcher-windows-x64.exe
	tar czvf switcher-linux-x64.tar.gz switcher-linux-x64 README.md LICENSE
	tar czvf switcher-darwin-x64.tar.gz switcher-darwin-x64 README.md LICENSE
	tar czvf switcher-windows-x64.tar.gz switcher-windows-x64.exe README.md LICENSE
