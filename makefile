GOOS = linux
GOARCH = amd64
OUTFILE = vless-filter-$(GOOS)-$(GOARCH)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTFILE) main.go

build-arm-darwin:
	$(MAKE) build GOOS=darwin GOARCH=arm64
build-amd64-darwin:
	$(MAKE) build GOOS=darwin GOARCH=amd64
build-arm-linux:
	$(MAKE) build GOOS=linux GOARCH=arm64
build-amd64-linux:
	$(MAKE) build GOOS=linux GOARCH=amd64
build-arm-windows:
	$(MAKE) build GOOS=windows GOARCH=arm64 OUTFILE=vless-filter-windows-arm64.exe
build-amd64-windows:
	$(MAKE) build GOOS=windows GOARCH=amd64 OUTFILE=vless-filter-windows-amd64.exe

build-all: build-arm-darwin build-amd64-darwin \
           build-arm-linux build-amd64-linux \
           build-arm-windows build-amd64-windows

clean:
	rm -f vless-filter-darwin-arm64 \
		vless-filter-darwin-amd64 \
		vless-filter-linux-arm64 \
		vless-filter-linux-amd64 \
		vless-filter-windows-arm64.exe \
		vless-filter-windows-amd64.exe