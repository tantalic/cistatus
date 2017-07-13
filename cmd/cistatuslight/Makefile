PKG := tantalic.com/cistatus
CMD := cmd/cistatuslight
SRC := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOVERSION := 1.8.3

.PHONY: run build linux-amd64 linux macos freebsd openbsd netbsd dragonfly windows docker-build docker-push binaries update-ca help
.DEFAULT_GOAL := help	

build:  ## Create binary suitable for use on a Raspberry Pi
	docker run --env GOOS=linux --env GOARCH=arm --env GOARM=6 --env CGO_ENABLED=0 --rm -v "$(SRC)":/go/src/$(PKG) -w /go/src/$(PKG)/$(CMD) golang:$(GOVERSION) go build -a -tags netgo -ldflags '-w' -o cistatuslight

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
