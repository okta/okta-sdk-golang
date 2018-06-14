# Some things this makefile could make use of:
#
# - test coverage target(s)
# - profiler target(s)
#

BIN             = rye
OUTPUT_DIR      = build
TMP_DIR        := .tmp
RELEASE_VER    := $(shell git rev-parse --short HEAD)
NAME            = default
COVERMODE       = atomic

TEST_PACKAGES      := $(shell go list ./... | grep -v vendor | grep -v fakes)

.PHONY: help
.DEFAULT_GOAL := help

run/examples:
	go run example/main.go

test: ##Run all tests
	go test ./...

test/codecov: ## Run all tests + open coverage report for all packages
	for PKG in $(TEST_PACKAGES); do \
		go test -covermode=$(COVERMODE) -coverprofile=profile.out $$PKG; \
		if [ -f profile.out ]; then\
			cat profile.out >> coverage.txt;\
			rm profile.out;\
		fi;\
	done
	$(RM) profile.out

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\/-]+:.*?## / {printf "\033[34m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | \
		sort | \
		grep -v '#'
