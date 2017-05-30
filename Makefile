.PHONY: test

all: Gomfile
	gom install

Gomfile:
	gom gen gomfile

SUCCESS := \e[1;32m
FAILURE := \e[1;31m
RESET   := \e[m

test:
	gom test -v ./src \
		| sed ''/PASS/s//$$(printf "$(SUCCESS)PASS$(RESET)")/'' \
		| sed ''/FAIL/s//$$(printf "$(FAILURE)FAIL$(RESET)")/''
