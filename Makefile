TEST_REGEXP ?= ""

.PHONY: test
test: # Run a specific test function
	go test -run $(TEST_REGEXP) ./...

.PHONY: tests
tests: # Run all tests
	go test ./...
