GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD)

build_main:
	$(GOBUILD) main.go

run:
	$(GORUN) *.go

test:
	$(GOTEST) -v ./animals
