cvgen-run:
	@go run cmd/cvgen/main.go $(ARGS)

# make cvgen-run ARGS="generate --input ./testdata/valid/cv.json --output /tmp/cv.html"

CVGEN_NAME := cvgen
CVGEN_BUILD_DIR ?= ./build

cvgen-build:
	mkdir -p $(CVGEN_BUILD_DIR)
	go build -o $(CVGEN_BUILD_DIR)/$(CVGEN_NAME) ./cmd/$(CVGEN_NAME)

cvgen-install: cvgen-build
	install -Dm755 $(CVGEN_BUILD_DIR)/$(CVGEN_NAME) /usr/local/bin/$(CVGEN_NAME)