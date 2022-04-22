COLOR_OK=\\x1b[0;32m
COLOR_NONE=\x1b[0m
COLOR_ERROR=\x1b[31;01m
COLOR_WARNING=\x1b[33;01m
COLOR_OKTA=\x1B[34;01m

SDKVERSION=$(shell grep -E -o 'API version: (0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?' ./client.go)
APIVERSION=$(shell grep -E -o '(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?' ./.swagger-codegen/VERSION)

help:
	@echo "$(COLOR_OKTA)   ____  _  _________             _                _                       $(COLOR_NONE)"
	@echo "$(COLOR_OKTA)  / __ \| |/ /__   __|/\         | |              | |                      $(COLOR_NONE)"
	@echo "$(COLOR_OKTA) | |  | | ' /   | |  /  \      __| | _____   _____| | ___  _ __   ___ _ __ $(COLOR_NONE)"
	@echo "$(COLOR_OKTA) | |  | |  <    | | / /\ \    / _  |/ _ \ \ / / _ \ |/ _ \| '_ \ / _ \ '__|$(COLOR_NONE)"
	@echo "$(COLOR_OKTA) | |__| | . \   | |/ ____ \  | (_| |  __/\ V /  __/ | (_) | |_) |  __/ |   $(COLOR_NONE)"
	@echo "$(COLOR_OKTA)  \____/|_|\_\  |_/_/    \_\  \__,_|\___| \_/ \___|_|\___/| .__/ \___|_|   $(COLOR_NONE)"
	@echo "$(COLOR_OKTA)                                                          | |              $(COLOR_NONE)"
	@echo "$(COLOR_OKTA)                                                          |_|              $(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_OK)Okta SDK Golang$(COLOR_NONE) version $(COLOR_WARNING)$(SDKVERSION)$(COLOR_NONE)"
	@echo "$(COLOR_OK)Swagger Codegen$(COLOR_NONE) version $(COLOR_WARNING)$(APIVERSION)$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Usage:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  make [command]$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Available commands:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  generate$(COLOR_NONE)       Generate the SDK with Swagger Codegen"
	@echo "$(COLOR_OK)  help$(COLOR_NONE)           Show this help message"
	@echo ""

generate:
	swagger-codegen generate -i /Users/bretterer/Code/OpenapiSpecV3/dist/management.yaml -l go
