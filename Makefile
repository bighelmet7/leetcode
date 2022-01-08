PKG_NAME ?= ""
TEST_REGEXP ?= ""

##@ Tests
.PHONY: test
test: # Run a specific test function
	go test -run $(TEST_REGEXP) ./...

.PHONY: tests
tests: # Run all tests
	go test ./...

##@ Helpers
.PHONY: coverage
coverage: # Get the coverage html
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: doc
doc: # Get documentation for one specific package
	go doc -u $(PKG_NAME)

##@ Cleaners
clean:
	rm -rf cover.html cover.out
