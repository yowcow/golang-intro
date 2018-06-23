.PHONY: test

all:
	which dep || go get -u -v github.com/golang/dep/cmd/dep
	dep ensure -v

test:
	go test ./src/...
