build:
	@go build -o build/app -ldflags="-s -w" -trimpath

build-race:
	@go build -o build/app-race -race -ldflags="-s -w" -trimpath
	
run: build
	@./build/app

%:
	@:

.PHONY: build build-race run
