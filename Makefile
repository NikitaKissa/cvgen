cvgen-run:
	@go run cmd/cvgen/main.go $(filter-out $@,$(MAKECMDGOALS))

%:
	@: