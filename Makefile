GO := GO111MODULE=on go

all:
	$(GO) get -u

test:
	$(GO) test ./src/...

clean:
	$(GO) clean -testcache || true

realclean: clean
	$(GO) clean -modcache || true

.PHONY: all test clean realclean
