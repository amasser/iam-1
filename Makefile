.PHONY: all tags clean test build install generate image release

REGISTRY_REPO = 940322424406.dkr.ecr.eu-central-1.amazonaws.com/iam

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Build flags
BUILD_DATE = $(shell date -u)
BUILD_HASH = $(shell git rev-parse --short HEAD)
BUILD_NUMBER ?= $(BUILD_NUMBER:)

# If we don't set the build number if defaults to dev
ifeq ($(BUILD_NUMBER), )
	BUILD_NUMBER := dev
endif

NOW := $(shell date -u '+%Y%m%d%I%M%S')

DOCKER := docker
GO := go
GO_ENV := $(shell $(GO) env GOOS GOARCH)
GOOS ?= $(word 1,$(GO_ENV))
GOARCH ?= $(word 2,$(GO_ENV))
GOFLAGS ?= $(GOFLAGS:)
ROOT_DIR := $(realpath .)

# GOOS/GOARCH of the build host, used to determine whether
# we're cross compiling or not
BUILDER_GOOS_GOARCH="$(GOOS)_$(GOARCH)"

# Add packages ./xxx/...
PKGS = $(shell $(GO) list ./cmd/... | grep -v /vendor/)

TAGS ?= "netgo"
BUILD_ENV = 
ENVFLAGS = CGOENABLED=1 $(BUILD_ENV)

ifneq ($(GOOS), darwin)
	EXTLDFLAGS = -extldflags "-lm -lstdc++ -static"
else
	EXTLDFLAGS = 
endif

GO_LINKER_FLAGS ?= --ldflags '$(EXTLDFLAGS) -s -w #\
	-X "github.com/maurofran/iam/pkg/version.BuildNumber=$(BUILD_NUMBER)" \
	-X "github.com/maurofran/iam/pkg/version.BuildDate=$(BUILD_DATE)" \
	-X "github.com/maurofran/iam/pkg/version.BuildHash=$(BUILD_HASH)"'

BIN_NAME := iamd

all: build

generate:
	@echo "$(OK_COLOR)==> Generating files via go generate...$(NO_COLOR)"
	@$(GO) generate $(GOFLAGS) $(PKGS)

build: generate
	@echo "$(OK_COLOR)==> Building binary ($(GOOS)/$(GOARCH))...$(NO_COLOR)"
	@echo @$(ENVFLAGS) $(GO) build -a -installsuffix cgo -tags $(TAGS) $(GOFLAGS) $(GO_LINKER_FLAGS) -o bin/$(GOOS)_$(GOARCH)/$(BIN_NAME) cmd/iamd/main.go
	@$(ENVFLAGS) $(GO) build -a -installsuffix cgo -tags $(TAGS) $(GOFLAGS) $(GO_LINKER_FLAGS) -o bin/$(GOOS)_$(GOARCH)/$(BIN_NAME) ./cmd/iamd/main.go

test:
	@echo "$(OK_COLOR)==> Running tests...$(NO_COLOR)"
	@$(GO) test $(GOFLAGS) $(PKGS)

install: build
	@echo "$(OK_COLOR)==> Installing packages into GOPATH...$(NO_COLOR)"
	@$(GO) install $(GOFLAGS) $(PKGS)

format:
	@echo "$(OK_COLOR)==> Formatting Code...$(NO_COLOR)"
	@$(GO) fmt $(GOFLAGS) $(PKGS)

vet:
	@echo "$(OK_COLOR)==> Running vet...$(NO_COLOR)"
	@$(GO) vet $(GOFLAGS) $(PKGS)

linter:
	@echo "$(OK_COLOR)==> Running linter...$(NO_COLOR)"
	@$(GO) lint $(GOFLAGS) $(PKGS)

setup:
	@echo "$(OK_COLOR)==> Installing required components...$(NO_COLOR)"
	@$(GO) get -u $(GOFLAGS) github.com/golang/dep/cmd/dep
	@dep ensure

clean:
	@echo "$(OK_COLOR)==> Cleaning...$(NO_COLOR)"
	@$(GO) clean -i ./...

run:
	@bin/$(GOOS)_$(GOARCH)/$(BIN_NAME) $(args)

image:
	@echo "$(OK_COLOR)==> Creating Docker Image...$(NO_COLOR)"
	@$(DOCKER) build . -t $(REGISTRY_REPO)

release:
	@echo "$(OK_COLOR)==> Pushing Docker Image to $(REGISTRY_REPO)...$(NO_COLOR)"
	@$(DOCKER) push $(REGISTRY_REPO)