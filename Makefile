GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOBENCH=$(GOTEST) -run=^$$ -benchmem -bench .
BINARY_NAME=gotherm.exe

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
bench:
	$(GOBENCH) ./...
.PHONY: clean
clean:
	- $(GOCLEAN)
	- rm -rf $(BINARY_NAME)
	- rm -rf dist/
run: build
	./$(BINARY_NAME) serve
tools: download
	cat tools.go | grep _ | awk -F'"' '{print $$2'} | xargs -tI % $(GOINSTALL) %
download:
	$(GOCMD) mod download
deps: download
	cd web && npm i && cd ..
release:
	cd web && npm run build && cd ..
	goreleaser --snapshot --rm-dist
