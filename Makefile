VERSION = 2.3.0-pre0

# ----------------------------------------------------------------------
# Test
# ----------------------------------------------------------------------
.PHONY: test
test: test-integration
	@echo "ran test task"

.PHONY: test-integration
test-integration:
	@echo "ran test-integration task"

# ----------------------------------------------------------------------
# Build examples
# ----------------------------------------------------------------------
EXAMPLES_DIR = examples

.PHONY: build-example-spin-http
build-example-spin-http:
	cd examples/spin-http && spin build

.PHONY: build-example-wasi-cli
build-example-wasi-cli:
	cd examples/wasi-cli && tinygo build -target=wasip2 -gc=leaking -no-debug -wit-package ../../wasi-cli/wit -wit-world wasi:cli/command -x -work -o main.wasm main.go

.PHONY: build-example-wasi-http
build-example-wasi-http:
	cd examples/wasi-http && tinygo build -target=wasip2 -gc=leaking -no-debug -wit-package ../../wasi-http/wit -wit-world wasi:http/proxy -x -work -o main.wasm main.go

.PHONY: build-examples
build-examples: build-example-spin-http

.PHONY: generate
generate: wit-bindgen-go
	@echo "removing generated code in internal/fermyon"
	rm -rf internal/fermyon
	@echo "removing generated code in internal/fermyon"
	rm -rf internal/wasi
	$(WIT_BINDGEN_GO) generate -p github.com/fermyon/spin-go-sdk/internal -o ./internal --exports ./wit
	$(WIT_BINDGEN_GO) generate -p github.com/fermyon/spin-go-sdk/internal -o ./internal --exports ./wasi-http/wit
	$(WIT_BINDGEN_GO) generate -p github.com/fermyon/spin-go-sdk/internal -o ./internal --exports ./wasi-cli/wit

# ----------------------------------------------------------------------
# Cleanup
# ----------------------------------------------------------------------
.PHONY: clean
clean:
	@rm -rf $(LOCALBIN)
	@echo "ran clean task"

# ----------------------------------------------------------------------
# Tools & Build Dependencies
# ----------------------------------------------------------------------

# TODO: what's the best way to get updated wit worlds? warg?

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

WIT_BINDGEN_GO ?= $(LOCALBIN)/wit-bindgen-go
.PHONY: wit-bindgen-go
wit-bindgen-go: $(WIT_BINDGEN_GO)
$(WIT_BINDGEN_GO): $(LOCALBIN)
	@if test -s $(WIT_BINDGEN_GO); then \
		echo "removing any existing installation of $(WIT_BINDGEN_GO)"; \
		rm $(WIT_BINDGEN_GO); \
	fi
	git submodule sync wasm-tools-go
	cd wasm-tools-go && GOBIN=$(LOCALBIN) go install ./...
