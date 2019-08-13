PROJECT_NAME := prometheus-demo
GOLANGCI_LINT_VER := v1.17.1

.PHONY: generate
generate:
	@go generate ./...

.PHONY: start
start:
	@echo "\033[0;32m» Open http://localhost:1337/metrics in your browser to view metrics.\033[0;39m"
	@go run cmd/prometheus-demo/main.go

.PHONY: lint
lint:
	@docker run --rm -w /src/github.com/martinsirbe/$(PROJECT_NAME) \
	    -v "$$PWD":/src/github.com/martinsirbe/$(PROJECT_NAME) \
	     golangci/golangci-lint:$(GOLANGCI_LINT_VER) golangci-lint run -v
