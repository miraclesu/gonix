# Targets:
# all: Build code.
# clean: Remove build artifacts.
# $FILENAME: Build just $FILENAME.
# build: Create build directory.

GC = go build
GCFLAGS =

BUILD_DIR = build
GO_FILES = $(wildcard *.go)
UTILS = $(patsubst %.go, $(BUILD_DIR)/%, $(GO_FILES))

.PHONY: all clean

all: $(UTILS)

$(BUILD_DIR)/%: %.go
	@mkdir -p $(BUILD_DIR)
	${GC} $(GCFLAGS) -o $@ $<

clean:
	rm -rf $(BUILD_DIR)

test: $(UTILS)
	go test ./tests
