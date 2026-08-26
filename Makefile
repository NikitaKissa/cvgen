cvgen-run:
	@go run cmd/cvgen/main.go $(ARGS)

# make cvgen-run ARGS="generate --input ./testdata/valid/cv.json --output /tmp/cv.html"

CVGEN_NAME := cvgen
CVGEN_BUILD_DIR ?= ./build

VERSION := $(shell git describe --tags --always --dirty)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -X 'github.com/NikitaKissa/cvgen/internal/cli.Version=$(VERSION)' \
           -X 'github.com/NikitaKissa/cvgen/internal/cli.Commit=$(COMMIT)' \
           -X 'github.com/NikitaKissa/cvgen/internal/cli.Date=$(DATE)'

cvgen-build:
	mkdir -p $(CVGEN_BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(CVGEN_BUILD_DIR)/$(CVGEN_NAME) ./cmd/$(CVGEN_NAME)

cvgen-install: cvgen-build
	install -Dm755 $(CVGEN_BUILD_DIR)/$(CVGEN_NAME) /usr/local/bin/$(CVGEN_NAME)