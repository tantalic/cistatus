PKG := tantalic.com/cistatus
CMD := cmd/cistatusanybar
SRC := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOVERSION := 1.8.3

.PHONY: build help
.DEFAULT_GOAL := help	

build: ## Create a macOS executable
	docker run --env GOOS=darwin --env GOARCH=amd64 --rm -v "$(SRC)":/go/src/$(PKG) -w /go/src/$(PKG)/$(CMD) golang:$(GOVERSION) go build -o cistatusanybar	

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
