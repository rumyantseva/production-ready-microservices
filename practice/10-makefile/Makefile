all: push

BUILDTAGS=

APP?=app
USERSPACE?=rumyantseva/production-ready-microservices/practice/10-makefile
RELEASE?=0.0.1
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux
PORT?=8000

CONTAINER_NAME?=${APP}

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

vendor: clean
	#go get -u github.com/Masterminds/glide \
	#&& glide install

build: vendor
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/version.RELEASE=${RELEASE} -X ${PROJECT}/version.COMMIT=${COMMIT} -X ${PROJECT}/version.REPO=${REPO_INFO}" \
		-o ${APP}

container: build
	docker build --pull -t $(APP):$(RELEASE) .

run: container
	docker run --name ${CONTAINER_NAME} -p ${PORT}:${PORT} \
		-e "PORT=${PORT}" \
		-d $(APP):$(RELEASE)

fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

lint:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint {{.Dir}}/..."{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

vet:
	@echo "+ $@"
	@go vet $(shell go list ${PROJECT}/... | grep -v vendor)

test: vendor fmt lint vet
	@echo "+ $@"
	@go test -v -race -tags "$(BUILDTAGS) cgo" $(shell go list ${PROJECT}/... | grep -v vendor)

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

clean:
	rm -f ${APP}
