# Go build variables
GOOS_WASM := js
GOARCH_WASM := wasm

# Detect host OS and architecture
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

# Set executable extension for Windows
ifeq ($(GOOS),windows)
	EXE_EXT := .exe
else
	EXE_EXT :=
endif

# Directories
BASELINE_WASM_SRC := ./cmd/baseline
SERVER_SRC := ./cmd/href_server_host
BUILDER_SRC := ./cmd/href_app
HREF_WASM_SRC := ./cmd/href_wasm
STATIC_OUT := ./static_output/demo1/web
SERVER_BIN := href_server_host
BUILDER_BIN := href_app
BASELINE_APP := baseline_app

.PHONY: all build build-wasm build-href-app generate-static run clean

all: build

build-baseline-wasm:
	@echo "Building Baseline WASM..."
	@mkdir -p web
	GOOS=$(GOOS_WASM) GOARCH=$(GOARCH_WASM) go build -o web/app.wasm $(BASELINE_WASM_SRC)

build-href-wasm:
	@echo "Building HREF WASM..."
	@mkdir -p $(STATIC_OUT)
	GOOS=$(GOOS_WASM) GOARCH=$(GOARCH_WASM) go build -o $(STATIC_OUT)/app.wasm $(HREF_WASM_SRC)

build: generate-static
	@echo "Building server..."
	go build -o $(SERVER_BIN)$(EXE_EXT) $(SERVER_SRC)

build-baseline-app:
	@echo "Building Baseline App..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BASELINE_APP)$(EXE_EXT) $(BASELINE_WASM_SRC)

build-href-app:
	@echo "Building HREF App..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILDER_BIN)$(EXE_EXT) $(BUILDER_SRC)

generate-static: build-baseline-wasm build-href-wasm build-href-app build-baseline-app
	@echo "Generating static site..."
	./$(BUILDER_BIN)$(EXE_EXT) -generate_static
	./$(BASELINE_APP)$(EXE_EXT) -generate_static

run: build
	@echo "Running server..."
	./$(SERVER_BIN)$(EXE_EXT)

clean:
	@echo "Cleaning..."
	rm -f $(STATIC_OUT)/app.wasm
	rm -f $(SERVER_BIN)
	rm -f $(SERVER_BIN)$(EXE_EXT)
	rm -f $(BUILDER_BIN)
	rm -f $(BUILDER_BIN)$(EXE_EXT)
	rm -f demo1_app$(EXE_EXT)