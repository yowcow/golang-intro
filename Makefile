DEP := $(GOPATH)/bin/dep
GO := GO111MODULE=on go

all: $(DEP)
	$(DEP) ensure -v
	$(GO) list -m

$(DEP):
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

test:
	$(GO) test ./src/...

.PHONY: all test
