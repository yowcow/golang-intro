GO := GO111MODULE=on go

all:

test:
	$(GO) test ./src/...

.PHONY: all test
