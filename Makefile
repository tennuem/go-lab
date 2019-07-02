test:
	go test -v -race ./pkg/...
bench:
	go test -bench=. ./pkg/... -benchmem > benchmem.log