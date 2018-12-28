GO := GO111MODULE=on go

WIRE_DIRS := ./mywire

all:

update:
	$(GO) get -u

wire:
	which wire || $(GO) get -u -v github.com/google/wire/cmd/wire
	wire $(WIRE_DIRS)

test:
	$(GO) test ./...

clean:
	$(GO) clean -testcache || true

realclean: clean
	$(GO) clean -modcache || true

.PHONY: all update wire test clean realclean
