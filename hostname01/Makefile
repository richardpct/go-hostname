.DEFAULT_GOAL := build
AWK           := awk
RM            := rm
GO            := go
GOFMT         := gofmt
PROGRAM       := hostname
BINARY        := go-$(PROGRAM)

.PHONY: help
help: ## Show help
	@echo "Usage: make TARGET\n"
	@echo "By default if TARGET is not provided the program will be compiled\n"
	@echo "Targets:"
	@$(AWK) -F ":.* ##" '/^[^#].*:.*##/{printf "%-13s%s\n", $$1, $$2}' \
	$(MAKEFILE_LIST) \
	| grep -v AWK

.PHONY: build
build: $(BINARY) ## Build

$(BINARY): $(PROGRAM).go
	$(GO) build -o $@

.PHONY: lint
lint: $(PROGRAM).go ## Check the syntax of the code
	$(GOFMT) -w $<

.PHONY: clean
clean: ## Delete the executable binary file
	$(RM) -fv $(BINARY)
