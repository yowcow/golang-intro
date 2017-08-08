.PHONY: test

all:
	rm -rf ./vendor ./Godeps
	godep save -v ./src/...

test:
	go test ./src/...
