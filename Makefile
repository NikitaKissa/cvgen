cvgen-run:
	@go run cmd/cvgen/main.go $(ARGS)

# make cvgen-run ARGS="generate --input ./testdata/valid/cv.json --output /tmp/cv.html"