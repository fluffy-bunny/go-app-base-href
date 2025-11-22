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
WASM_SRC := ./cmd/demo1
SERVER_SRC := ./cmd/server
BUILDER_SRC := ./cmd/builder
STATIC_OUT := ./static_output/demo1/web
SERVER_BIN := server
BUILDER_BIN := builder

.PHONY: all build build-wasm build-builder generate-static run clean

all: build

build-wasm-baseline:
	@echo "Building Baseline WASM..."
	@mkdir -p web
	GOOS=$(GOOS_WASM) GOARCH=$(GOARCH_WASM) go build -o web/app.wasm $(WASM_SRC)


build: generate-static
	@echo "Building server..."
	go build -o $(SERVER_BIN) $(SERVER_SRC)

build-wasm:
	@echo "Building WASM..."
	@mkdir -p $(STATIC_OUT)
	GOOS=$(GOOS_WASM) GOARCH=$(GOARCH_WASM) go build -o $(STATIC_OUT)/app.wasm $(BUILDER_SRC)

build-demo1-app:
	@echo "Building Demo1 App..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o demo1_app$(EXE_EXT) $(WASM_SRC)

build-builder:
	@echo "Building builder..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILDER_BIN)$(EXE_EXT) $(BUILDER_SRC)

generate-static: build-wasm-baseline build-wasm build-builder build-demo1-app
	@echo "Generating static site..."
	./$(BUILDER_BIN)$(EXE_EXT) -generate_static
	./demo1_app$(EXE_EXT) -generate_static

run: build
	@echo "Running server..."
	./$(SERVER_BIN)

clean:
	@echo "Cleaning..."
	rm -f $(STATIC_OUT)/app.wasm
	rm -f $(SERVER_BIN)
	rm -f $(SERVER_BIN)$(EXE_EXT)
	rm -f $(BUILDER_BIN)
	rm -f $(BUILDER_BIN)$(EXE_EXT)
	rm -f demo1_app$(EXE_EXT)