GIT_COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date -u +"%FT%T%z")
APP_VERSION = "${GIT_COMMIT}_${BUILD_DATE}"
BINARY = amy
LDFLAGS = -ldflags "-X main.Version=${APP_VERSION}"

.PHONY: build
build:
	go build ${LDFLAGS} -o ${BINARY} main.go

.PHONY: setup
setup:
	go get -u github.com/govend/govend
	govend -v

.PHONY: run
run: build
	./${BINARY}
