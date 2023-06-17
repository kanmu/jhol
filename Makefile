.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/jhol

.PHONY: vet
vet:
	go vet -composites=false -structtag=false ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -f jhol
