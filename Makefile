# text reset
NO_COLOR=\033[0m
# green
OK_COLOR=\033[32;01m
# red
ERROR_COLOR=\033[31;01m
# cyan
WARN_COLOR=\033[36;01m
# yellow
ATTN_COLOR=\033[33;01m

ROOT_DIR := $(git rev-parse --show-toplevel)
BIN_DIR  := ./bin

LINTER := $(BIN_DIR)/golangci-lint
TESTRUNNER := $(GOPATH)/bin/gotestsum

GOOS :=
ifeq ($(OS),Windows_NT)
	GOOS = windows
else 
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		GOOS = darwin
	endif
endif
GOARCH ?= amd64

VERSION:=`git describe --tags --dirty 2>/dev/null`
COMMIT :=`git rev-parse --short HEAD 2>/dev/null`
DATE   :=`date "+%FT%T%z"`

LDFLAGS := -ldflags "-w -s -X github.com/gertd/go-pluralize/cmd/pluralize/version.version=${VERSION} -X github.com/gertd/go-pluralize/cmd/pluralize/version.date=${DATE} -X github.com/gertd/go-pluralize/cmd/pluralize/version.commit=${COMMIT}"

BINARY := pluralize
PLATFORMS := windows linux darwin
OS = $(word 1, $@)

.PHONY: all
all: build test lint

deps:
	@echo -e "$(ATTN_COLOR)==> download dependencies $(NO_COLOR)"
	@GO111MODULE=on go mod download

.PHONY: build
build: deps
	@echo -e "$(ATTN_COLOR)==> build GOOS=$(GOOS) GOARCH=$(GOARCH) VERSION=$(VERSION) COMMIT=$(COMMIT) DATE=$(DATE) $(NO_COLOR)"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=on go build $(LDFLAGS) -o $(BIN_DIR)/$(BINARY) ./cmd/pluralize

$(TESTRUNNER):
	@echo -e "$(ATTN_COLOR)==> get gotestsum test runner  $(NO_COLOR)"
	@go get -u gotest.tools/gotestsum 

.PHONY: test
test: $(TESTRUNNER)
	@echo -e "$(ATTN_COLOR)==> test $(NO_COLOR)"
	@gotestsum --format short-verbose

$(LINTER):
	@echo -e "$(ATTN_COLOR)==> get  $(NO_COLOR)"
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.15.0
 
.PHONY: lint
lint: $(LINTER)
	@echo -e "$(ATTN_COLOR)==> lint $(NO_COLOR)"
	@$(LINTER) run --enable-all
	@echo -e "$(NO_COLOR)\c"

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@echo -e "$(ATTN_COLOR)==> release GOOS=$(OS) GOARCH=$(GOARCH) release/$(BINARY)-$(OS)-$(GOARCH) $(NO_COLOR)"
	@mkdir -p release
	@GOOS=$(OS) GOARCH=$(GOARCH) GO111MODULE=on go build $(LDFLAGS) -o release/$(BINARY)-$(OS)-$(GOARCH)$(if $(findstring $(OS),windows),".exe","")  ./cmd/pluralize

.PHONY: release
release: windows linux darwin

.PHONY: install
install:
	@echo -e "$(ATTN_COLOR)==> install $(NO_COLOR)"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=on go install $(LDFLAGS) ./cmd/pluralize
