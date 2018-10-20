GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gothermostat.exe

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
.PHONY: clean
clean:
	- $(GOCLEAN)
	- rm -rf $(BINARY_NAME)
	- rm -rf dist/
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	cd web && npm i && cd ..
	$(GOGET) github.com/smartystreets/goconvey/convey
	$(GOGET) github.com/goreleaser/goreleaser
release:
	cd web && npm run build && cd ..
	goreleaser --snapshot
