.PHONY: test

all:
	rm -rf ./vendor
	dep ensure -update

test:
	go test ./src/...
