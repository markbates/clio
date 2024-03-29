TAGS ?= ""
GO_BIN ?= "go"

install:
	pkger
	$(GO_BIN) install -tags ${TAGS} -v ./cmd/clio
	make tidy

tidy:
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
else
	echo skipping go mod tidy
endif

deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...
	make tidy

build:
	pkger
	$(GO_BIN) build -v .
	make tidy

test:
	pkger
	$(GO_BIN) test -cover -tags ${TAGS} ./...
	make tidy

ci-deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...

ci-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

cov:
	$(GO_BIN) test -coverprofile cover.out -tags ${TAGS} ./...
	go tool cover -html cover.out
	make tidy

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --enable-all
	make tidy

update:
ifeq ($(GO111MODULE),on)
	rm go.*
	$(GO_BIN) mod init
	$(GO_BIN) mod tidy
else
	$(GO_BIN) get -u -tags ${TAGS}
endif
	make test
	make install
	make tidy

release-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...
	make tidy

release:
	pkger
	make tidy
	release -y -f version.go --skip-packr
	make tidy



