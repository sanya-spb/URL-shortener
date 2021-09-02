PROJECT?=github.com/sanya-spb/URL-shortener
PROJECTNAME=$(shell basename "$(PROJECT)")

GOOS?=linux
GOARCH?=amd64

CGO_ENABLED=1
# EXE_FILE=url-shortener

RELEASE := $(shell git tag -l | tail -1 | grep -E "v.+"|| echo devel)
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COPYRIGHT := "sanya-spb"

## build-test: Build test-env
build-test:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=$(CGO_ENABLED) go build \
		-ldflags "-s -w \
		-X ${PROJECT}/pkg/version.version=${RELEASE} \
		-X ${PROJECT}/pkg/version.commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.buildTime=${BUILD_TIME} \
		-X ${PROJECT}/pkg/version.copyright=${COPYRIGHT}" \
		-o ./cmd/test-env/test-env ./cmd/test-env/

## build: Build url-shortener
build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=$(CGO_ENABLED) go build \
		-ldflags "-s -w \
		-X ${PROJECT}/pkg/version.version=${RELEASE} \
		-X ${PROJECT}/pkg/version.commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.buildTime=${BUILD_TIME} \
		-X ${PROJECT}/pkg/version.copyright=${COPYRIGHT}" \
		-o ./cmd/url-shortener/url-shortener ./cmd/url-shortener/

## check: Run linters
check:
	golangci-lint -c ./.golangci.yml run

## run: Run url-shortener
run:
	go run ./cmd/url-shortener/ -config ./data/conf/config.yaml

## clean: Clean build files
clean:
	go clean
	rm -v ./cmd/test-env/test-env 2> /dev/null || true
	rm -v ./cmd/url-shortener/url-shortener 2> /dev/null || true
	rm -v ./data/logs/*.log 2> /dev/null || true

## test: Run unit test
test:
	go test -v -short ${PROJECT}/cmd/url-shortener/

## integration: Run integration test
integration:
	# go test -v -run Integration ${PROJECT}/cmd/url-shortener/

## help: Show this
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
