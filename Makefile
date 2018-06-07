COLOR_OK=\\x1b[0;32m
COLOR_NONE=\x1b[0m
COLOR_ERROR=\x1b[31;01m
COLOR_WARNING=\x1b[33;05m
COLOR_OKTA=\x1B[34;01m

help:
	@echo "$(COLOR_OK)Okta SDK for Golang$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Usage:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  make [command]$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Available commands:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  build                   Build the Okta Golang SDK$(COLOR_NONE)"
	@echo "$(COLOR_OK)  help                    Shows this help$(COLOR_NONE)"
	@echo "$(COLOR_OK)  pull-spec               Pull down the most recent released version of the spec$(COLOR_NONE)"
	@echo "$(COLOR_WARNING)test$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:all                Run all tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:integration        Run only integration tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:unit               Run only unit tests$(COLOR_NONE)"

build:
	@echo "$(COLOR_OKTA)Building SDK...$(COLOR_NONE)"
	@echo "TODO"

pull-spec:
	@echo "$(COLOR_OKTA)Pulling in latest spec...$(COLOR_NONE)"
	rm openapi/spec.json
	git clone https://github.com/okta/openapi spec-raw
	cp spec-raw/dist/spec.json openapi/spec.json
	rm -fr spec-raw

test\:all:
	@echo "$(COLOR_OKTA)Running all tests...$(COLOR_NONE)"
	@echo "TODO"

test\:integration:
	@echo "$(COLOR_OKTA)Running integration tests...$(COLOR_NONE)"
	@echo "TODO"

test\:unit:
	@echo "$(COLOR_OK)Running unit tests...$(COLOR_NONE)"
	@echo "TODO"
