
GO      	  :=  go
PROMU   	  := $(GOPATH)/bin/promu
PROMU_VERSION := 0.13.0
PREFIX  	  ?= $(shell pwd)
BIN_DIR       ?= $(shell pwd)
PKGS	 	  = $(shell $(GO) list ./... | grep -v /vendor/)
ARCH          = $(shell uname -m)
GOOS          = $(shell uname -s | tr A-Z a-z)

ifeq ($(ARCH), aarch64)
	GOARCH=arm64
endif
ifeq ($(ARCH), x86_64)
	GOARCH=amd64
endif
ifeq ($(ARCH), x86)
	GOARCH=386
endif
ifeq ($(ARCH), armv7l)
	GOARCH=arm
endif

all: test build

test:
	@echo ">> running tests"
	go test -short $(PKGS)

build: promu
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

tarballs: promu
	@echo ">> building release tarballs"
	@$(PROMU) crossbuild tarballs
	@echo ">> calculating release checksums"
	@$(PROMU) checksum $(BIN_DIR)/.tarballs

crossbuild: promu
	@echo ">> cross-building binaries"
	@$(PROMU) crossbuild

release: promu
	@echo ">> create release"
	@$(PROMU) release --verbose --timeout=120s --retry=30 $(BIN_DIR)/.tarballs

promu:
	$(GO) install github.com/prometheus/promu@v$(PROMU_VERSION)
