.PHONY: test vet fmt-check

test:
	go test ./... -v

vet:
	go vet ./...

fmt-check:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)
