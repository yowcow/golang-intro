.PHONY: test

all: Gomfile
	gom install

Gomfile:
	gom gen gomfile

SUCCESS := \033[1;32m
FAILURE := \033[1;31m
RESET   := \033[m

test:
	gom test -v ./src \
		| sed ''/PASS/s//$$(printf "$(SUCCESS)PASS$(RESET)")/'' \
		| sed ''/FAIL/s//$$(printf "$(FAILURE)FAIL$(RESET)")/''
