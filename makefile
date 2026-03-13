build-arm-darwin:
	GOOS=darwin GOARCH=arm64 go build -o vless-filter-darwin-arm64 main.go
build-amd64-darwin:
	GOOS=darwin GOARCH=amd64 go build -o vless-filter-darwin-amd64 main.go
build-arm-linux:
	GOOS=linux GOARCH=arm64 go build -o vless-filter-linux-arm
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o vless-filter-linux-amd64 main.go
build-arm-windows:
	GOOS=windows GOARCH=arm64 go build -o vless-filter-windows-arm64.exe main.go
build-amd64-windows:
	GOOS=windows GOARCH=amd64 go build -o vless-filter-windows-amd64.exe main.go
build-all:
	$(MAKE) build-arm-darwin
	$(MAKE) build-amd64-darwin
	$(MAKE) build-arm-linux
	$(MAKE) build-amd64-linux
	$(MAKE) build-arm-windows
	$(MAKE) build-amd64-windows
clean:
	rm -f vless-filter-darwin-arm64 \
		vless-filter-darwin-amd64 \
		vless-filter-linux-arm \
		vless-filter-linux-amd64 \
		vless-filter-windows-arm64.exe \
		vless-filter-windows-amd64.exe