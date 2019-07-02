vendor:
	go mod vendor 
test: vendor
	go test -v -race ./pkg/...
bench:
	go test -bench=. ./pkg/... -benchmem > benchmem.log