GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOBENCH=$(GOTEST) -run=^$$ -benchmem -bench .
VERSION ?= $(shell git describe --tags --abbrev 2> /dev/null || echo v0.0.0-dev)
MAJOR ?= $(word 1,$(subst ., ,$(VERSION)))
MINOR ?= $(word 2,$(subst ., ,$(VERSION)))
BINARY_NAME=gotherm.exe

.PHONY: all build bench clean test tools deps download run dockerize-web dockerize-publish-web
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
dockerize-web:
	docker build -t marcofranssen/gotherm-web web
	docker rmi $$(docker images -qf dangling=true)
docker-publish-web:
	docker tag marcofranssen/gotherm-web:latest marcofranssen/gotherm-web:$(VERSION)
	docker tag marcofranssen/gotherm-web:latest marcofranssen/gotherm-web:$(MAJOR).$(MINOR)
	docker tag marcofranssen/gotherm-web:latest marcofranssen/gotherm-web:$(MAJOR)
release:
	cd web && npm run build && cd ..
	goreleaser --snapshot --rm-dist
