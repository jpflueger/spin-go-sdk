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

.PHONY: build-examples
build-examples: 
	@echo "ran build-examples task"

.PHONY: generate
generate: wit-bindgen-go
	@echo "generating http-trigger world"
	$(WIT_BINDGEN_GO) generate -w fermyon:spin/http-trigger@2.0.0 -p github.com/fermyon/spin-go-sdk/generated -o ./generated --exports ./wit
	$(WIT_BINDGEN_GO) generate -w wasi:http/proxy@0.2.0 -p github.com/fermyon/spin-go-sdk/generated -o ./generated --exports ./wasi-http/wit
	$(WIT_BINDGEN_GO) generate -w wasi:cli/command@0.2.0 -p github.com/fermyon/spin-go-sdk/generated -o ./generated --exports ./wasi-cli/wit
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
