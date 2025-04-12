.PHONY: all info build run clean

# Default target
all: clean templ-gen build run

# Print information about available commands
info:
	$(info This Makefile helps you manage your projects.)
	$(info )
	$(info Available commands:)
	$(info - build:  Build the Golang project.)
	$(info - all:  Run all commands (Build ,run).)
	$(info - templ-gen Run templ generate) 
	$(info Usage: make <command>)

# Build the Golang project
build:
	@echo "=== Building Golang Project ==="
	@go build -o go-ota main.go

# Build the Golang project
run:
	@echo "=== Running Golang Project ==="
	@./go-ota

# Clean build artifacts
clean:
	@echo "=== Cleaning build artifacts ==="
	@rm -f go-ota

# Clean build artifacts
templ-gen:
	@echo "=== Generating templ files ==="
	@templ generate