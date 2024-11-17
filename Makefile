
.PHONY: build
build:
	@rm -rf bin
	@go build -o bin/devservices ./cmd/devservices
	@echo "Success Build"