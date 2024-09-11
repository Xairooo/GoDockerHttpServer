SHELL := /bin/bash

# Read .env file
include .env
export $(shell sed 's/=.*//' .env)

# Default target
.PHONY: all
all: build

# Build Docker image
.PHONY: build
build:
	@echo "Building Docker image..."
	docker build $(shell while IFS= read -r line || [[ -n "$$line" ]]; do \
		[[ "$$line" =~ ^#.*$$ ]] && continue; \
		[[ -z "$$line" ]] && continue; \
		line=$$(echo "$$line" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$$//'); \
		echo "--build-arg $$line"; \
	done < .env) -t go_docker_http_server .

# Run Docker container
.PHONY: run
run:
	@echo "Running Docker container..."
	docker run -p $(PORT):$(PORT) --env-file .env go_docker_http_server

# Clean up
.PHONY: clean
clean:
	@echo "Cleaning up..."
	docker rmi go_docker_http_server

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build  - Build Docker image"
	@echo "  run    - Run Docker container"
	@echo "  clean  - Remove Docker image"
	@echo "  help   - Show this help message"