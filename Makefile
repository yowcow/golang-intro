.PHONY: test

all: Gomfile
	gom install

Gomfile:
	gom gen gomfile

test:
	gom test -v ./src
