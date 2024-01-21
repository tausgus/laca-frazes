GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BINARY_NAME=server
BINARY_PATH=bin/$(BINARY_NAME)

all: build

build:
	$(GOBUILD) -o $(BINARY_PATH) ./cmd/server/

clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

.PHONY: all build clean
