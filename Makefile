DEP := $(GOPATH)/bin/dep
GO := GO111MODULE=on go

all: $(DEP)
	$(DEP) ensure -v
	$(GO) list -m

$(DEP):
	$(GO) get -u -v github.com/golang/dep/cmd/dep

test:
	$(GO) test ./src/...

.PHONY: all test
