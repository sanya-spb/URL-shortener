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

## build: Build application
build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=$(CGO_ENABLED) go build \
		-ldflags "-s -w \
		-X ${PROJECT}/pkg/version.version=${RELEASE} \
		-X ${PROJECT}/pkg/version.commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.buildTime=${BUILD_TIME} \
		-X ${PROJECT}/pkg/version.copyright=${COPYRIGHT}" \
		-o ./cmd/test-env/test-env ./cmd/test-env/

## check: Run linters
check:
	golangci-lint -c ./golangci-lint.yaml run

## run: Run application
run: check test build
	./cmd/test-env/test-env -config ./data/test-env/config.toml

## clean: Clean build files
clean:
	go clean
	rm ./cmd/test-env/test-env

## test: Run unit test
test:
	go test -v -short ${PROJECT}/cmd/test-env/

## integration: Run integration test
integration:
	# go test -v -run Integration ${PROJECT}/utils/fdouble/

## help: Show this
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
