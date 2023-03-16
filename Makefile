# Variables
GOARCH := wasm
GOOS := js
WASM_DIR := wasm
PUBLIC_DIR := public

# Targets
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "Usage: make <target>"
	@echo "Available targets:"
	@echo "  <wasm>    Build wasm file from ./wasm/<wasm>"
	@echo "  clean     Remove all build artifacts"
.PHONY: help

%:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o ./$(PUBLIC_DIR)/$@.wasm ./$(WASM_DIR)/$@/

clean: ## Remove all build artifacts
	rm -rf ./$(PUBLIC_DIR)/*.wasm

# Special targets
.SUFFIXES:
.DELETE_ON_ERROR:
.SECONDARY:
