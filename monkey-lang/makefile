.PHONY: test
test:
	go test -timeout 5s ./...

.PHONY: clean
clean:
	go clean --testcache

.PHONY: clean_test
clean_test: clean test
	exit 0