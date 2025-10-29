COLOR_OK=\\x1b[0;32m
COLOR_NONE=\x1b[0m
COLOR_ERROR=\x1b[31;01m
COLOR_WARNING=\x1b[33;05m
COLOR_OKTA=\x1B[34;01m
GOLANGCI_LINT := golangci-lint
GOLANGCI_LINT_VERSION := v2.5.0

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
	cd openapi && yarn generator
	@echo "$(COLOR_OK)Running goimports and gofumpt on generated files...$(COLOR_NONE)"
	@make import
	@make fmt
	@echo "$(COLOR_OKTA)Done!$(COLOR_NONE)"

pull-spec:
	@echo "$(COLOR_OKTA)Pulling in latest spec...$(COLOR_NONE)"
	cd openapi && npm install
# Overwrite the spec.json file from npm if a git branch is specified by env variable.
ifneq ($(origin OPENAPI_SPEC_BRANCH),undefined)
	rm -f openapi/spec.json
	git clone --branch $(OPENAPI_SPEC_BRANCH) https://github.com/okta/okta-management-openapi-spec spec-raw
	cp spec-raw/dist/spec.json openapi/spec.json
	cp spec-raw/dist/spec.json openapi/node_modules/@okta/openapi/dist/spec.json
	rm -fr spec-raw
endif

.PHONY: fmt
fmt: check-golangci-lint # Format the code using `golangci-lint`
	@$(GOLANGCI_LINT) fmt

.PHONY: import
import: # Run goimports on all Go files
	@goimports -w .

check-golangci-lint:
	@which $(GOLANGCI_LINT) > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

test:
	go test -cover -coverpkg=./okta -failfast -race ./okta ./okta/test -test.v

test\:all:
	go test -cover -coverpkg=./okta -failfast -race ./okta ./okta/test -test.v

test\:integration:
	go test -cover -coverpkg=./okta -failfast -race ./okta/test -test.v

test\:unit:
	go test -cover -coverpkg=./okta -failfast -race ./okta -test.v

generate:
	npx @openapitools/openapi-generator-cli generate -c ./.generator/config.yaml -i .generator/okta-management-APIs-oasv3-noEnums-inheritance.yaml --skip-validate-spec

