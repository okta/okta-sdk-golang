COLOR_OK=\\x1b[0;32m
COLOR_NONE=\x1b[0m
COLOR_ERROR=\x1b[31;01m
COLOR_WARNING=\x1b[33;05m
COLOR_OKTA=\x1B[34;01m
GOFMT := gofumpt
GOIMPORTS := goimports

ifndef OPENAPI_SPEC_BRANCH
override OPENAPI_SPEC_BRANCH = master
endif

help:
	@echo "$(COLOR_OK)Okta SDK for Golang$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Usage:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  make [command]$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Available commands:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  build                   Clean and build the Okta Golang SDK generated files$(COLOR_NONE)"
	@echo "$(COLOR_OK)  clean-files             Deletes all generated files$(COLOR_NONE)"
	@echo "$(COLOR_OK)  generate-files          Generates files based around spec$(COLOR_NONE)"
	@echo "$(COLOR_OK)  pull-spec               Pull down the most recent released version of the spec$(COLOR_NONE)"
	@echo "$(COLOR_WARNING)test$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:all                Run all tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:integration        Run only integration tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:unit               Run only unit tests$(COLOR_NONE)"

build:
	@echo "$(COLOR_OKTA)Building SDK...$(COLOR_NONE)"
	make clean-files
	make generate-files
	make test:all

clean-files:
	@echo "$(COLOR_OKTA)Cleaning Up Old Generated Files...$(COLOR_NONE)"
	cd openapi && yarn cleanFiles

generate-files:
	@echo "$(COLOR_OKTA)Generating SDK Files...$(COLOR_NONE)"
	cd openapi && npm install && yarn generator
	@echo "$(COLOR_OK)Running goimports and gofumpt on generated files...$(COLOR_NONE)"
	@make import
	@make fmt
	@echo "$(COLOR_OKTA)Done!$(COLOR_NONE)"

pull-spec:
	@echo "$(COLOR_OKTA)Pulling in latest spec...$(COLOR_NONE)"
	rm -f openapi/spec.json
	git clone --branch $(OPENAPI_SPEC_BRANCH) https://github.com/okta/okta-management-openapi-spec spec-raw
	cp spec-raw/dist/spec.json openapi/spec.json
	cp spec-raw/dist/spec.json openapi/node_modules/@okta/openapi/dist/spec.json
	rm -fr spec-raw

test:
	make test:all

test\:all:
	@echo "$(COLOR_OKTA)Running all tests...$(COLOR_NONE)"
	@make test:unit
	@make test:integration

test\:integration:
	@echo "$(COLOR_OKTA)Running integration tests...$(COLOR_NONE)"
	go test -race ./tests/integration -test.v

test\:unit:
	@echo "$(COLOR_OK)Running unit tests...$(COLOR_NONE)"
	go test -race ./tests/unit -test.v

.PHONY: fmt
fmt: check-fmt # Format the code
	@$(GOFMT) -l -w $$(find . -name '*.go' |grep -v vendor) > /dev/null

check-fmt:
	@which $(GOFMT) > /dev/null || GO111MODULE=on go get mvdan.cc/gofumpt

.PHONY: import
import: check-goimports
	@$(GOIMPORTS) -w $$(find . -path ./vendor -prune -o -name '*.go' -print) > /dev/null

check-goimports:
	@which $(GOIMPORTS) > /dev/null || GO111MODULE=on go get golang.org/x/tools/cmd/goimports
