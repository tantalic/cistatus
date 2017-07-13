PKG := tantalic.com/cistatus
CMD := cmd/cistatusserver
SRC := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
GOVERSION := 1.8.3

DOCKER_IMAGE := "tantalic/cistatusserver"

COMMIT = $(strip $(shell git rev-parse --short HEAD))
VERSION := $(strip $(shell git describe --always --dirty))

.PHONY: linux-amd64 docker-build docker-push update-ca help
.DEFAULT_GOAL := help	

linux-amd64:
	docker run --env GOOS=linux --env GOARCH=amd64 --env CGO_ENABLED=0 --rm -v "$(SRC)":/go/src/$(PKG) -w /go/src/$(PKG)/$(CMD) golang:$(GOVERSION) go build -a -tags netgo -ldflags '-w' -o cistatusserver-linux_amd64

docker-image: linux-amd64 ## Build a docker image
	docker build \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		-t $(DOCKER_IMAGE):$(VERSION) .

docker-push: ## Push the docker image to DockerHub
	docker push $(DOCKER_IMAGE):$(VERSION)

update-ca: ## Download the latest CA roots
	curl --time-cond ca-certificates.crt -o ca-certificates.crt https://curl.haxx.se/ca/cacert.pem

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
