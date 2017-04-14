PKG := tantalic.com/cistatus
CMD := cmd/cistatusanybar
SRC := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOVERSION := 1.8.1

.PHONY: build help
.DEFAULT_GOAL := help	

all: cistatusserver cistatusanybar cistatuslight ## Build all artifacts (see below)

cistatusserver: ## Build docker image (and linux executable) for cistatusserver
	cd cmd/cistatusserver && make docker-image

cistatusanybar: ## Build cistatusanybar macOS executable
	cd cmd/cistatusanybar && make build

cistatuslight: ## Build cistatuslight executable (ARM)
	cd cmd/cistatuslight && make build

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
