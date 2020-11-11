.PHONY: test
test:
	go test -v -race ./...

.PHONY: bench
bench:
	go test -bench=. ./pkg/... -benchmem > benchmem.log