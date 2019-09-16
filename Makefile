WIRE_DIRS := ./mywire

all: update wire

update:
	go mod tidy

wire:
	which wire || go get -u -v github.com/google/wire/cmd/wire
	wire $(WIRE_DIRS)

test:
	go test ./...

clean:
	go clean -testcache || true

realclean: clean
	go clean -modcache || true

.PHONY: all update wire test clean realclean
