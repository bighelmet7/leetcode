.PHONY: test
test: # Run tests for a given package
	@read -p "Test Function pattern: " TEST_REGEXP; \
	go test -run $$TEST_REGEXP ./...

.PHONY: tests
tests: # Run all tests
	go test ./...
