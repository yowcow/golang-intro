.PHONY: test

all:
	rm -rf ./vendor
	dep ensure

test:
	go test ./src/...
