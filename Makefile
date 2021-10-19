
GO      	  :=  go
PROMU   	  := $(GOPATH)/bin/promu
PROMU_VERSION := 0.12.0
pkgs	 	  = $(shell $(GO) list ./... | grep -v /vendor/)
PREFIX  	  ?= $(shell pwd)
BIN_DIR       ?= $(shell pwd)


all: test build

test:
	@echo ">> running tests"
	go test -short $(pkgs)

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
	@$(PROMU) release --verbose --timeout=60s --retry=30 $(BIN_DIR)/.tarballs

promu:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
	        GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
	        $(GO) get -u github.com/prometheus/promu@v$(PROMU_VERSION)
			
