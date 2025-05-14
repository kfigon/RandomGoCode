.PHONY: test
test:
	go test -timeout 5s ./...

.PHONY: clean
clean:
	go clean --testcache

.PHONY: ctest
ctest: clean test
	exit 0