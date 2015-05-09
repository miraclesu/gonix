# Targets:
# all: Build code.
# clean: Remove build artifacts.
# $FILENAME: Build just $FILENAME.
# build: Create build directory.

GC = go build
GCFLAGS =

BUILD_DIR = build
GO_FILES = $(wildcard *.go)
UTILS = $(basename $(GO_FILES))

.PHONY: all clean

all: $(UTILS)

%: %.go $(BUILD_DIR)
	${GC} $(GCFLAGS) -o $(BUILD_DIR)/$@ $<

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

testing:
	cd tests && go test
