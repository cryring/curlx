GO=CGO_ENABLED=1 go

default: build

build:
	$(GO) build -o curlx ./cmd/main.go

clean:
	@rm -f curlx
	@echo "clean done"